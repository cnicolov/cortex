load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["driver.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/daemon/graphdriver:go_default_library",
        "//vendor/github.com/docker/docker/pkg/chrootarchive:go_default_library",
        "//vendor/github.com/docker/docker/pkg/idtools:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/label:go_default_library",
    ],
)
