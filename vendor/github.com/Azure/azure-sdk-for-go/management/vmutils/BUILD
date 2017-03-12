load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "configurationset.go",
        "datadisks.go",
        "deployment.go",
        "extensions.go",
        "network.go",
        "rolesize.go",
        "rolestate.go",
        "vmutils.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/azure-sdk-for-go/management:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/location:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachine:go_default_library",
        "//vendor/github.com/Azure/azure-sdk-for-go/management/virtualmachinedisk:go_default_library",
    ],
)
