load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "api.go",
        "conv.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/gorilla/mux:go_default_library",
        "//vendor/github.com/openzipkin/zipkin-go-opentracing/_thrift/gen-go/zipkincore:go_default_library",
        "//vendor/github.com/weaveworks-experiments/loki/pkg/storage:go_default_library",
    ],
)
