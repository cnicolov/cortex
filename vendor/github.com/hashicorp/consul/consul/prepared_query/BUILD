load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "template.go",
        "walk.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/consul/consul/structs:go_default_library",
        "//vendor/github.com/hashicorp/hil:go_default_library",
        "//vendor/github.com/hashicorp/hil/ast:go_default_library",
        "//vendor/github.com/mitchellh/copystructure:go_default_library",
    ],
)
