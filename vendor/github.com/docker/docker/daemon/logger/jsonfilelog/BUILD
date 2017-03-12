load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "jsonfilelog.go",
        "read.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/docker/daemon/logger:go_default_library",
        "//vendor/github.com/docker/docker/daemon/logger/loggerutils:go_default_library",
        "//vendor/github.com/docker/docker/pkg/filenotify:go_default_library",
        "//vendor/github.com/docker/docker/pkg/ioutils:go_default_library",
        "//vendor/github.com/docker/docker/pkg/jsonlog:go_default_library",
        "//vendor/github.com/docker/docker/pkg/tailfile:go_default_library",
        "//vendor/github.com/docker/go-units:go_default_library",
        "//vendor/github.com/fsnotify/fsnotify:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
