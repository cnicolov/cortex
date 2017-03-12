load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["manifestlist.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution:go_default_library",
        "//vendor/github.com/docker/distribution/digest:go_default_library",
        "//vendor/github.com/docker/distribution/manifest:go_default_library",
    ],
)
