load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "compare.go",
        "config.go",
        "config_unix.go",
        "errors.go",
        "hostconfig.go",
        "hostconfig_unix.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/errors:go_default_library",
        "//vendor/github.com/docker/docker/api/types/container:go_default_library",
        "//vendor/github.com/docker/docker/api/types/network:go_default_library",
        "//vendor/github.com/docker/docker/pkg/sysinfo:go_default_library",
        "//vendor/github.com/docker/docker/volume:go_default_library",
    ],
)
