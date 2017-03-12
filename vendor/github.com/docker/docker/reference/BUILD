load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "reference.go",
        "store.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution/digest:go_default_library",
        "//vendor/github.com/docker/distribution/reference:go_default_library",
        "//vendor/github.com/docker/docker/image/v1:go_default_library",
        "//vendor/github.com/docker/docker/pkg/ioutils:go_default_library",
    ],
)
