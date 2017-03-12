load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["local.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/docker/pkg/idtools:go_default_library",
        "//vendor/github.com/docker/docker/pkg/mount:go_default_library",
        "//vendor/github.com/docker/docker/utils:go_default_library",
        "//vendor/github.com/docker/docker/volume:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)
