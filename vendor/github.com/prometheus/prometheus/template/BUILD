load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["template.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/prometheus/common/model:go_default_library",
        "//vendor/github.com/prometheus/prometheus/promql:go_default_library",
        "//vendor/github.com/prometheus/prometheus/util/strutil:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
