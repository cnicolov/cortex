load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["register_vfs.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/daemon/graphdriver/vfs:go_default_library"],
)
