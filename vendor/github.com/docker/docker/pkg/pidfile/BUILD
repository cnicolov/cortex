load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "pidfile.go",
        "pidfile_darwin.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/pkg/system:go_default_library"],
)
