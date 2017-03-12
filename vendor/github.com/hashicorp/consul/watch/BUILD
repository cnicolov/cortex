load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "funcs.go",
        "plan.go",
        "watch.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/hashicorp/consul/api:go_default_library"],
)
