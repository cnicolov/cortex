load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["print.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution/reference:go_default_library",
        "//vendor/github.com/docker/docker/api/types/swarm:go_default_library",
        "//vendor/github.com/docker/docker/cli/command:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/idresolver:go_default_library",
        "//vendor/github.com/docker/docker/pkg/stringid:go_default_library",
        "//vendor/github.com/docker/go-units:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
