load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["salsa20.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/crypto/salsa20/salsa:go_default_library"],
)
