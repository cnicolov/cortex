load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/pierrec/lz4:go_default_library"],
)

go_binary(
    name = "lz4c",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
