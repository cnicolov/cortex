load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "cobra.go",
        "error.go",
        "required.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/spf13/cobra:go_default_library"],
)
