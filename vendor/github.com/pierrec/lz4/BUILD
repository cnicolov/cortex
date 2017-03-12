load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "block.go",
        "lz4.go",
        "reader.go",
        "writer.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/pierrec/xxHash/xxHash32:go_default_library"],
)
