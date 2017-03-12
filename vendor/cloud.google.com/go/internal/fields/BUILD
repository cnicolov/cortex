load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "fields.go",
        "fold.go",
    ],
    visibility = ["//vendor/cloud.google.com/go:__subpackages__"],
    deps = ["//vendor/cloud.google.com/go/internal/atomiccache:go_default_library"],
)
