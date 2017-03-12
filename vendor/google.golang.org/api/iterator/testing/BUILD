load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["testing.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/google.golang.org/api/iterator:go_default_library"],
)
