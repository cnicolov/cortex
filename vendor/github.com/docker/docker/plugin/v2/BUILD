load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "plugin.go",
        "plugin_unsupported.go",
        "settable.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution/digest:go_default_library",
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/docker/docker/pkg/plugingetter:go_default_library",
        "//vendor/github.com/docker/docker/pkg/plugins:go_default_library",
        "//vendor/github.com/opencontainers/runtime-spec/specs-go:go_default_library",
    ],
)
