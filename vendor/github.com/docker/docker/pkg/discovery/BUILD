load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "backends.go",
        "discovery.go",
        "entry.go",
        "generator.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/Sirupsen/logrus:go_default_library"],
)
