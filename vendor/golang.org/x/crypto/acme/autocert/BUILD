load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "autocert.go",
        "cache.go",
        "renewal.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/crypto/acme:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
