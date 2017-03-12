load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "progress.go",
        "progressreader.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/time/rate:go_default_library"],
)
