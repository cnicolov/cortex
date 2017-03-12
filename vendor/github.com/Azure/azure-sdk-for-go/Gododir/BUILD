load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["gen.go"],
    visibility = ["//visibility:private"],
    deps = ["//vendor/gopkg.in/godo.v2:go_default_library"],
)

go_binary(
    name = "Gododir",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
