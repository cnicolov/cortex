load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["commands.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/cli/command:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/checkpoint:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/container:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/image:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/network:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/node:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/plugin:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/registry:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/secret:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/service:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/stack:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/swarm:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/system:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/volume:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
    ],
)
