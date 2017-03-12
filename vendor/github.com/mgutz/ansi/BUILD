load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "ansi.go",
        "doc.go",
        "print.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/mattn/go-colorable:go_default_library"],
)
