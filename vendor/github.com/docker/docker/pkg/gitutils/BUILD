load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["gitutils.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/symlink:go_default_library",
        "//vendor/github.com/docker/docker/pkg/urlutil:go_default_library",
    ],
)
