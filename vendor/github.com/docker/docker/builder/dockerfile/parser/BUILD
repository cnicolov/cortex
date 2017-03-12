load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "line_parsers.go",
        "parser.go",
        "utils.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/docker/docker/builder/dockerfile/command:go_default_library"],
)
