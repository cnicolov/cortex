load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["restartmanager.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/api/types/container:go_default_library"],
)
