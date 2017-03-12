load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "entities.go",
        "resourceextensions.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/azure-sdk-for-go/management:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachinedisk:go_default_library",
    ],
)
