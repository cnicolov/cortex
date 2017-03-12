load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["file.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/mattes/migrate/migrate/direction:go_default_library"],
)
