load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["gccgoexportdata.go"],
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/tools/go/internal/gccgoimporter:go_default_library"],
)
