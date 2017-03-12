load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "httpclient.go",
        "httpserver.go",
        "implementation.go",
        "service.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/opentracing/opentracing-go:go_default_library",
        "//vendor/github.com/openzipkin/zipkin-go-opentracing/examples/cli_with_2_services/svc2:go_default_library",
        "//vendor/github.com/openzipkin/zipkin-go-opentracing/examples/middleware:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
