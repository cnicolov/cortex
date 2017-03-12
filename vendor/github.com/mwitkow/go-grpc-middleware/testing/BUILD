load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "interceptor_suite.go",
        "pingservice.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/mwitkow/go-grpc-middleware/testing/testproto:go_default_library",
        "//vendor/github.com/stretchr/testify/require:go_default_library",
        "//vendor/github.com/stretchr/testify/suite:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/credentials:go_default_library",
    ],
)
