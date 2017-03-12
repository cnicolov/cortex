load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["prune.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/cli/command:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/container:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/image:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/network:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/volume:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
    ],
)
