load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "errors.go",
        "http.go",
        "operations.go",
        "publishSettings.go",
        "util.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/golang.org/x/crypto/pkcs12:go_default_library"],
)
