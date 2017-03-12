load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["echo.go"],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "webhook",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
