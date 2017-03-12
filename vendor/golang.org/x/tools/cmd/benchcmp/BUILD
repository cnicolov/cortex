load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "benchcmp.go",
        "compare.go",
        "doc.go",
    ],
    visibility = ["//visibility:private"],
    deps = ["//vendor/golang.org/x/tools/benchmark/parse:go_default_library"],
)

go_binary(
    name = "benchcmp",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
