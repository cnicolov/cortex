load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "gated_writer.go",
        "log_levels.go",
        "log_writer.go",
        "logger.go",
        "syslog.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/go-syslog:go_default_library",
        "//vendor/github.com/hashicorp/logutils:go_default_library",
        "//vendor/github.com/mitchellh/cli:go_default_library",
    ],
)
