load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["utils.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/integration/cmd:go_default_library",
        "//vendor/github.com/docker/docker/pkg/stringutils:go_default_library",
        "//vendor/github.com/docker/docker/pkg/system:go_default_library",
    ],
)
