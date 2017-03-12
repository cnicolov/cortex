load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "callgraph.go",
        "util.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/tools/go/ssa:go_default_library"],
)
