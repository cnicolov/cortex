load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["mem.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/prometheus/alertmanager/provider:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/types:go_default_library",
        "//vendor/github.com/prometheus/common/model:go_default_library",
    ],
)
