load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["interpolation.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/cli/compose/template:go_default_library",
        "//vendor/github.com/docker/docker/cli/compose/types:go_default_library",
    ],
)
