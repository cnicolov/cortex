load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["windows.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/go-ansiterm:go_default_library",
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
    ],
)
