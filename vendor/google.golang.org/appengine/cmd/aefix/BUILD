load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "ae.go",
        "fix.go",
        "main.go",
        "typecheck.go",
    ],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "aefix",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
