package distributor

import (
	"fmt"
	"testing"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/common/model"
	"github.com/prometheus/prometheus/storage/metric"
	"github.com/stretchr/testify/assert"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/weaveworks/common/user"
	"github.com/weaveworks/cortex/ingester/api"
	"github.com/weaveworks/cortex/ring"
)

// mockRing doesn't do any consistent hashing, just returns same ingesters for every query.
type mockRing struct {
	prometheus.Counter
	ingesters []*ring.IngesterDesc
}

func (r mockRing) Get(key uint32, n int, op ring.Operation) ([]*ring.IngesterDesc, error) {
	return r.ingesters[:n], nil
}

func (r mockRing) BatchGet(keys []uint32, n int, op ring.Operation) ([][]*ring.IngesterDesc, error) {
	result := [][]*ring.IngesterDesc{}
	for i := 0; i < len(keys); i++ {
		result = append(result, r.ingesters[:n])
	}
	return result, nil
}

func (r mockRing) GetAll() []*ring.IngesterDesc {
	return r.ingesters
}

type mockIngester struct {
	happy bool
}

func (i mockIngester) Push(ctx context.Context, in *api.WriteRequest, opts ...grpc.CallOption) (*api.WriteResponse, error) {
	if !i.happy {
		return nil, fmt.Errorf("Fail")
	}
	return &api.WriteResponse{}, nil
}

func (i mockIngester) Query(ctx context.Context, in *api.QueryRequest, opts ...grpc.CallOption) (*api.QueryResponse, error) {
	if !i.happy {
		return nil, fmt.Errorf("Fail")
	}
	return &api.QueryResponse{
		Timeseries: []api.TimeSeries{
			{
				Labels: []api.LabelPair{
					{
						Name:  []byte("__name__"),
						Value: []byte("foo"),
					},
				},
				Samples: []api.Sample{
					{
						Value:       0,
						TimestampMs: 0,
					},
					{
						Value:       1,
						TimestampMs: 1,
					},
				},
			},
		},
	}, nil
}

func (i mockIngester) LabelValues(ctx context.Context, in *api.LabelValuesRequest, opts ...grpc.CallOption) (*api.LabelValuesResponse, error) {
	return nil, nil
}

func (i mockIngester) UserStats(ctx context.Context, in *api.UserStatsRequest, opts ...grpc.CallOption) (*api.UserStatsResponse, error) {
	return nil, nil
}

func (i mockIngester) MetricsForLabelMatchers(ctx context.Context, in *api.MetricsForLabelMatchersRequest, opts ...grpc.CallOption) (*api.MetricsForLabelMatchersResponse, error) {
	return nil, nil
}

func TestDistributorPush(t *testing.T) {
	ctx := user.Inject(context.Background(), "user")
	for i, tc := range []struct {
		ingesters        []mockIngester
		samples          int
		expectedResponse *api.WriteResponse
		expectedError    error
	}{
		// A push of no samples shouldn't block or return error, even if ingesters are sad
		{
			ingesters:        []mockIngester{{}, {}, {}},
			expectedResponse: &api.WriteResponse{},
		},

		// A push to 3 happy ingesters should succeed
		{
			samples:          10,
			ingesters:        []mockIngester{{true}, {true}, {true}},
			expectedResponse: &api.WriteResponse{},
		},

		// A push to 2 happy ingesters should succeed
		{
			samples:          10,
			ingesters:        []mockIngester{{}, {true}, {true}},
			expectedResponse: &api.WriteResponse{},
		},

		// A push to 1 happy ingesters should fail
		{
			samples:       10,
			ingesters:     []mockIngester{{}, {}, {true}},
			expectedError: fmt.Errorf("Fail"),
		},

		// A push to 0 happy ingesters should fail
		{
			samples:       10,
			ingesters:     []mockIngester{{}, {}, {}},
			expectedError: fmt.Errorf("Fail"),
		},
	} {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			ingesterDescs := []*ring.IngesterDesc{}
			ingesters := map[string]mockIngester{}
			for i, ingester := range tc.ingesters {
				addr := fmt.Sprintf("%d", i)
				ingesterDescs = append(ingesterDescs, &ring.IngesterDesc{
					Addr:      addr,
					Timestamp: time.Now().Unix(),
				})
				ingesters[addr] = ingester
			}

			ring := mockRing{
				Counter: prometheus.NewCounter(prometheus.CounterOpts{
					Name: "foo",
				}),
				ingesters: ingesterDescs,
			}

			d, err := New(Config{
				ReplicationFactor:   3,
				HeartbeatTimeout:    1 * time.Minute,
				RemoteTimeout:       1 * time.Minute,
				ClientCleanupPeriod: 1 * time.Minute,
				IngestionRateLimit:  10000,
				IngestionBurstSize:  10000,

				ingesterClientFactory: func(addr string) api.IngesterClient {
					return ingesters[addr]
				},
			}, ring)
			if err != nil {
				t.Fatal(err)
			}
			defer d.Stop()

			request := &api.WriteRequest{}
			for i := 0; i < tc.samples; i++ {
				ts := api.TimeSeries{
					Labels: []api.LabelPair{
						{[]byte("__name__"), []byte("foo")},
						{[]byte("bar"), []byte("baz")},
						{[]byte("sample"), []byte(fmt.Sprintf("%d", i))},
					},
				}
				ts.Samples = []api.Sample{
					{
						Value:       float64(i),
						TimestampMs: int64(i),
					},
				}
				request.Timeseries = append(request.Timeseries, ts)
			}
			response, err := d.Push(ctx, request)
			assert.Equal(t, tc.expectedResponse, response, "Wrong response")
			assert.Equal(t, tc.expectedError, err, "Wrong error")
		})
	}
}

func TestDistributorQuery(t *testing.T) {
	ctx := user.Inject(context.Background(), "user")

	expectedResponse := func(start, end int) model.Matrix {
		result := model.Matrix{
			&model.SampleStream{
				Metric: model.Metric{"__name__": "foo"},
			},
		}
		for i := start; i < end; i++ {
			result[0].Values = append(result[0].Values,
				model.SamplePair{
					Value:     model.SampleValue(i),
					Timestamp: model.Time(i),
				},
			)
		}
		return result
	}

	for i, tc := range []struct {
		ingesters        []mockIngester
		expectedResponse model.Matrix
		expectedError    error
	}{
		// A query to 3 happy ingesters should succeed
		{
			ingesters:        []mockIngester{{true}, {true}, {true}},
			expectedResponse: expectedResponse(0, 2),
		},

		// A query to 2 happy ingesters should succeed
		{
			ingesters:        []mockIngester{{}, {true}, {true}},
			expectedResponse: expectedResponse(0, 2),
		},

		// A query to 1 happy ingesters should fail
		{
			ingesters:     []mockIngester{{}, {}, {true}},
			expectedError: fmt.Errorf("Fail"),
		},

		// A query to 0 happy ingesters should succeed
		{
			ingesters:     []mockIngester{{}, {}, {}},
			expectedError: fmt.Errorf("Fail"),
		},
	} {
		t.Run(fmt.Sprintf("[%d]", i), func(t *testing.T) {
			ingesterDescs := []*ring.IngesterDesc{}
			ingesters := map[string]mockIngester{}
			for i, ingester := range tc.ingesters {
				addr := fmt.Sprintf("%d", i)
				ingesterDescs = append(ingesterDescs, &ring.IngesterDesc{
					Addr:      addr,
					Timestamp: time.Now().Unix(),
				})
				ingesters[addr] = ingester
			}

			ring := mockRing{
				Counter: prometheus.NewCounter(prometheus.CounterOpts{
					Name: "foo",
				}),
				ingesters: ingesterDescs,
			}

			d, err := New(Config{
				ReplicationFactor:   3,
				HeartbeatTimeout:    1 * time.Minute,
				RemoteTimeout:       1 * time.Minute,
				ClientCleanupPeriod: 1 * time.Minute,
				IngestionRateLimit:  10000,
				IngestionBurstSize:  10000,

				ingesterClientFactory: func(addr string) api.IngesterClient {
					return ingesters[addr]
				},
			}, ring)
			if err != nil {
				t.Fatal(err)
			}
			defer d.Stop()

			matcher, err := metric.NewLabelMatcher(metric.Equal, model.LabelName("__name__"), model.LabelValue("foo"))
			if err != nil {
				t.Fatal(err)
			}
			response, err := d.Query(ctx, 0, 10, matcher)
			assert.Equal(t, tc.expectedResponse, response, "Wrong response")
			assert.Equal(t, tc.expectedError, err, "Wrong error")
		})
	}
}
