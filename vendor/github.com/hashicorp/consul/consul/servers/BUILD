load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["manager.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/consul/consul/agent:go_default_library",
        "//vendor/github.com/hashicorp/consul/lib:go_default_library",
    ],
)
