load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "kernel.go",
        "kernel_darwin.go",
        "uname_unsupported.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/mattn/go-shellwords:go_default_library"],
)
