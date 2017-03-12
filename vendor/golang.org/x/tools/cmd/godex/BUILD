load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "gc.go",
        "gccgo.go",
        "godex.go",
        "isAlias18.go",
        "print.go",
        "source.go",
        "writetype.go",
    ],
    visibility = ["//visibility:private"],
)

go_binary(
    name = "godex",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
