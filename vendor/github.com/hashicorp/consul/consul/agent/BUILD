load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["server.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/hashicorp/serf/serf:go_default_library"],
)
