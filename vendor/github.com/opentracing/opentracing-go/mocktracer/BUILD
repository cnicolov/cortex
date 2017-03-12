load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "mocklogrecord.go",
        "mockspan.go",
        "mocktracer.go",
        "propagation.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/opentracing/opentracing-go:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go/ext:go_default_library",
        "//vendor/github.com/opentracing/opentracing-go/log:go_default_library",
    ],
)
