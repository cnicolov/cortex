load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "godoc.go",
        "talks.go",
        "tip.go",
    ],
    visibility = ["//visibility:private"],
    deps = ["//vendor/golang.org/x/crypto/acme/autocert:go_default_library"],
)

go_binary(
    name = "tip",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
