load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/opentracing/opentracing-go:go_default_library",
        "//vendor/github.com/openzipkin/zipkin-go-opentracing:go_default_library",
        "//vendor/github.com/openzipkin/zipkin-go-opentracing/examples/cli_with_2_services/svc2:go_default_library",
    ],
)

go_binary(
    name = "cmd",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
