load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "credentials.go",
        "default_store.go",
        "default_store_darwin.go",
        "file_store.go",
        "native_store.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker-credential-helpers/client:go_default_library",
        "//vendor/github.com/docker/docker-credential-helpers/credentials:go_default_library",
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/docker/docker/cliconfig/configfile:go_default_library",
        "//vendor/github.com/docker/docker/registry:go_default_library",
    ],
)
