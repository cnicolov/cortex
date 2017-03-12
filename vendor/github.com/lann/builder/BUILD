load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "builder.go",
        "reflect.go",
        "registry.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/lann/ps:go_default_library"],
)
