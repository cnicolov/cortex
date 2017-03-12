load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "handler.go",
        "main.go",
        "server.go",
    ],
    visibility = ["//visibility:private"],
    deps = ["//vendor/git.apache.org/thrift.git/lib/go/thrift:go_default_library"],
)

go_binary(
    name = "src",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
