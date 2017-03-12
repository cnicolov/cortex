load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/pierrec/xxHash/xxHash32:go_default_library",
        "//vendor/github.com/pierrec/xxHash/xxHash64:go_default_library",
    ],
)

go_binary(
    name = "xxhsum",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
