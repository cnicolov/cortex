load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "access.go",
        "htpasswd.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution/context:go_default_library",
        "//vendor/github.com/docker/distribution/registry/auth:go_default_library",
        "//vendor/golang.org/x/crypto/bcrypt:go_default_library",
    ],
)
