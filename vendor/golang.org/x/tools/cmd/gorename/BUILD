load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/golang.org/x/tools/go/buildutil:go_default_library",
        "//vendor/golang.org/x/tools/refactor/rename:go_default_library",
    ],
)

go_binary(
    name = "gorename",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
