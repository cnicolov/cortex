load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "test_response_writer.go",
        "test_round_tripper.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/stretchr/testify/mock:go_default_library"],
)
