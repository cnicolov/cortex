load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "commands.go",
        "main.go",
        "version.go",
    ],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/hashicorp/serf/cmd/serf/command:go_default_library",
        "//vendor/github.com/hashicorp/serf/cmd/serf/command/agent:go_default_library",
        "//vendor/github.com/mitchellh/cli:go_default_library",
    ],
)

go_binary(
    name = "serf",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
