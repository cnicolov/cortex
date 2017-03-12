load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["helpers.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/Azure/go-autorest/autorest/azure:go_default_library"],
)
