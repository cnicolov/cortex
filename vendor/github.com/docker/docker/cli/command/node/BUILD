load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cmd.go",
        "demote.go",
        "inspect.go",
        "list.go",
        "opts.go",
        "promote.go",
        "ps.go",
        "remove.go",
        "update.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/docker/docker/api/types/swarm:go_default_library",
        "//vendor/github.com/docker/docker/cli:go_default_library",
        "//vendor/github.com/docker/docker/cli/command:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/idresolver:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/inspect:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/task:go_default_library",
        "//vendor/github.com/docker/docker/client:go_default_library",
        "//vendor/github.com/docker/docker/opts:go_default_library",
        "//vendor/github.com/docker/docker/pkg/ioutils:go_default_library",
        "//vendor/github.com/docker/docker/runconfig/opts:go_default_library",
        "//vendor/github.com/docker/go-units:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/github.com/spf13/pflag:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
