load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/golang/snappy:go_default_library",
        "//vendor/github.com/influxdb/influxdb/client:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/github.com/prometheus/common/log:go_default_library",
        "//vendor/github.com/prometheus/common/model:go_default_library",
        "//vendor/github.com/prometheus/prometheus/documentation/examples/remote_storage/remote_storage_bridge/graphite:go_default_library",
        "//vendor/github.com/prometheus/prometheus/documentation/examples/remote_storage/remote_storage_bridge/influxdb:go_default_library",
        "//vendor/github.com/prometheus/prometheus/documentation/examples/remote_storage/remote_storage_bridge/opentsdb:go_default_library",
        "//vendor/github.com/prometheus/prometheus/storage/remote:go_default_library",
    ],
)

go_binary(
    name = "remote_storage_bridge",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
