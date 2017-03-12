load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "extractor.go",
        "oauth2.go",
        "request.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/dgrijalva/jwt-go:go_default_library"],
)
