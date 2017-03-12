load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "archive.go",
        "snapshot.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/hashicorp/raft:go_default_library"],
)
