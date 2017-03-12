load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "eg.go",
        "match.go",
        "rewrite.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/tools/go/ast/astutil:go_default_library"],
)
