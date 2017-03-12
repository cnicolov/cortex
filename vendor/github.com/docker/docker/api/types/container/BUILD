load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "container_create.go",
        "container_update.go",
        "container_wait.go",
        "host_config.go",
        "hostconfig_unix.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types/blkiodev:go_default_library",
        "//vendor/github.com/docker/docker/api/types/mount:go_default_library",
        "//vendor/github.com/docker/docker/api/types/strslice:go_default_library",
        "//vendor/github.com/docker/go-connections/nat:go_default_library",
        "//vendor/github.com/docker/go-units:go_default_library",
    ],
)
