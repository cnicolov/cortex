load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["benchmark.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/benchmark/grpc_testing:go_default_library",
        "//vendor/google.golang.org/grpc/grpclog:go_default_library",
    ],
)
