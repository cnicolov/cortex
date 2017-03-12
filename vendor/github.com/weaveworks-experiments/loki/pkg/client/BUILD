load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "collector.go",
        "tracer.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/apache/thrift/lib/go/thrift:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go:go_default_library",
        "//vendor/github.com/openzipkin/zipkin-go-opentracing:go_default_library",
        "//vendor/github.com/openzipkin/zipkin-go-opentracing/_thrift/gen-go/zipkincore:go_default_library",
    ],
)
