load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "discovery.go",
        "discovery_unix.go",
        "errors.go",
        "plugins.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/docker/pkg/plugins/transport:go_default_library",
        "//vendor/github.com/docker/go-connections/sockets:go_default_library",
        "//vendor/github.com/docker/go-connections/tlsconfig:go_default_library",
    ],
)
