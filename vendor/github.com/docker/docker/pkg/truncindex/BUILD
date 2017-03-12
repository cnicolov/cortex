load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["truncindex.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/tchap/go-patricia/patricia:go_default_library"],
)
