load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["lz4.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/pierrec/lz4:go_default_library"],
)
