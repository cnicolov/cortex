load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "archive.go",
        "archive_unix.go",
        "chroot_unix.go",
        "diff.go",
        "diff_unix.go",
        "init_unix.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/pkg/archive:go_default_library",
        "//vendor/github.com/docker/docker/pkg/idtools:go_default_library",
        "//vendor/github.com/docker/docker/pkg/reexec:go_default_library",
        "//vendor/github.com/docker/docker/pkg/system:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/system:go_default_library",
    ],
)
