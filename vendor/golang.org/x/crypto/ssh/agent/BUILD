load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "forward.go",
        "keyring.go",
        "server.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/crypto/ed25519:go_default_library",
        "//vendor/golang.org/x/crypto/ssh:go_default_library",
    ],
)
