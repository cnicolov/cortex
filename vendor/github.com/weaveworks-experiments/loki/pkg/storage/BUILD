load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "boltdb.go",
        "immutable_block.go",
        "in_memory.go",
        "merge.go",
        "mutable_block.go",
        "storage.go",
        "trace.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/apache/thrift/lib/go/thrift:go_default_library",
        "//vendor/github.com/boltdb/bolt:go_default_library",
        "//vendor/github.com/openzipkin/zipkin-go-opentracing/_thrift/gen-go/zipkincore:go_default_library",
        "//vendor/github.com/sburnett/lexicographic-tuples:go_default_library",
        "//vendor/github.com/willf/bloom:go_default_library",
    ],
)
