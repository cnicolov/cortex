load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["generate.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/docker/docker/cli/command:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/commands:go_default_library",
        "//vendor/github.com/docker/docker/pkg/term:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/spf13/cobra/doc:go_default_library",
    ],
)

go_binary(
    name = "man",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
