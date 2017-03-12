load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["syslog.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/RackSec/srslog:go_default_library",
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/docker/daemon/logger:go_default_library",
        "//vendor/github.com/docker/docker/daemon/logger/loggerutils:go_default_library",
        "//vendor/github.com/docker/docker/pkg/urlutil:go_default_library",
        "//vendor/github.com/docker/go-connections/tlsconfig:go_default_library",
    ],
)
