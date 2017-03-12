load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["stathat.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/rcrowley/go-metrics:go_default_library",
        "//vendor/github.com/stathat/go:go_default_library",
    ],
)
