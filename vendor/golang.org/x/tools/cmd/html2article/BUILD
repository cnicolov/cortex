load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["conv.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/golang.org/x/net/html:go_default_library",
        "//vendor/golang.org/x/net/html/atom:go_default_library",
    ],
)

go_binary(
    name = "html2article",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
