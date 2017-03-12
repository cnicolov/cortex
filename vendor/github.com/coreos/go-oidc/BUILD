load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "jose.go",
        "jwks.go",
        "oidc.go",
        "verify.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/pquerna/cachecontrol:go_default_library",
        "//vendor/golang.org/x/oauth2:go_default_library",
        "//vendor/gopkg.in/square/go-jose.v2:go_default_library",
    ],
)
