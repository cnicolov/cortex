load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "base64.go",
        "bcrypt.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/crypto/blowfish:go_default_library"],
)
