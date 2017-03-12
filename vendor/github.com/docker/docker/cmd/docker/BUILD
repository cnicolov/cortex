load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "daemon_none.go",
        "docker.go",
    ],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/docker/api/types/versions:go_default_library",
        "//vendor/github.com/docker/docker/cli:go_default_library",
        "//vendor/github.com/docker/docker/cli/command:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/commands:go_default_library",
        "//vendor/github.com/docker/docker/cli/flags:go_default_library",
        "//vendor/github.com/docker/docker/cliconfig:go_default_library",
        "//vendor/github.com/docker/docker/dockerversion:go_default_library",
        "//vendor/github.com/docker/docker/pkg/term:go_default_library",
        "//vendor/github.com/docker/docker/utils:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
    ],
)

go_binary(
    name = "docker",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
