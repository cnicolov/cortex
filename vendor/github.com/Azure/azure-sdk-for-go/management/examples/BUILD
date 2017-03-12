load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["example.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/Azure/azure-sdk-for-go/management:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/hostedservice:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/storageservice:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachine:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/vmutils:go_default_library",
    ],
)

go_binary(
    name = "examples",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
