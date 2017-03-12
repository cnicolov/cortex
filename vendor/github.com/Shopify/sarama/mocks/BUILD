load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "async_producer.go",
        "consumer.go",
        "mocks.go",
        "sync_producer.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/Shopify/sarama:go_default_library"],
)
