load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["create.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/Azure/azure-sdk-for-go/arm/examples/helpers:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/arm/storage:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/azure:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/to:go_default_library",
    ],
)

go_binary(
    name = "create",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
