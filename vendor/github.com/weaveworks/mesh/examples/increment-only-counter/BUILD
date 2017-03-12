load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "main.go",
        "peer.go",
        "state.go",
    ],
    visibility = ["//visibility:private"],
    deps = ["//vendor/github.com/weaveworks/mesh:go_default_library"],
)

go_binary(
    name = "increment-only-counter",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
