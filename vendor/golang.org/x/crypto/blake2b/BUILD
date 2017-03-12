load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "blake2b.go",
        "blake2bAVX2_amd64.go",
        "blake2bAVX2_amd64.s",
        "blake2b_amd64.s",
        "blake2b_generic.go",
    ],
    visibility = ["//visibility:public"],
)
