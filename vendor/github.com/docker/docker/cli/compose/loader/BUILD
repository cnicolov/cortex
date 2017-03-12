load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["loader.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/cli/compose/interpolation:go_default_library",
        "//vendor/github.com/docker/docker/cli/compose/schema:go_default_library",
        "//vendor/github.com/docker/docker/cli/compose/types:go_default_library",
        "//vendor/github.com/docker/docker/runconfig/opts:go_default_library",
        "//vendor/github.com/docker/go-units:go_default_library",
        "//vendor/github.com/mattn/go-shellwords:go_default_library",
        "//vendor/github.com/mitchellh/mapstructure:go_default_library",
        "//vendor/gopkg.in/yaml.v2:go_default_library",
    ],
)
