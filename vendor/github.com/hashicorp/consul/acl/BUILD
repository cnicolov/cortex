load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "acl.go",
        "cache.go",
        "policy.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/armon/go-radix:go_default_library",
        "//vendor/github.com/hashicorp/golang-lru:go_default_library",
        "//vendor/github.com/hashicorp/hcl:go_default_library",
    ],
)
