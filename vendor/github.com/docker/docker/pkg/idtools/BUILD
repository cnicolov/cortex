load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "idtools.go",
        "idtools_unix.go",
        "usergroupadd_unsupported.go",
        "utils_unix.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/system:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/user:go_default_library",
    ],
)
