load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "internal.go",
        "match.go",
        "tables.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/text/language:go_default_library"],
)
