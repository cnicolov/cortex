load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["kafka-console-partitionconsumer.go"],
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/Shopify/sarama:go_default_library"],
)

go_binary(
    name = "kafka-console-partitionconsumer",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
