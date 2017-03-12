load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "const.go",
        "rpc_client.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/go-msgpack/codec:go_default_library",
        "//vendor/github.com/hashicorp/logutils:go_default_library",
        "//vendor/github.com/hashicorp/serf/coordinate:go_default_library",
        "//vendor/github.com/hashicorp/serf/serf:go_default_library",
    ],
)
