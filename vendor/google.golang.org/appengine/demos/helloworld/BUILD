load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["helloworld.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/google.golang.org/appengine:go_default_library",
        "//vendor/google.golang.org/appengine/log:go_default_library",
    ],
)

go_binary(
    name = "helloworld",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
