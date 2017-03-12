load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "login.go",
        "logout.go",
        "search.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/docker/docker/api/types/registry:go_default_library",
        "//vendor/github.com/docker/docker/cli:go_default_library",
        "//vendor/github.com/docker/docker/cli/command:go_default_library",
        "//vendor/github.com/docker/docker/opts:go_default_library",
        "//vendor/github.com/docker/docker/pkg/stringutils:go_default_library",
        "//vendor/github.com/docker/docker/registry:go_default_library",
        "//vendor/github.com/spf13/cobra:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
