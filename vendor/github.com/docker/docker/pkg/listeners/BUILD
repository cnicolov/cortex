load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["listeners_unix.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/coreos/go-systemd/activation:go_default_library",
        "//vendor/github.com/docker/go-connections/sockets:go_default_library",
    ],
)
