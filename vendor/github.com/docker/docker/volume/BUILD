load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "validate.go",
        "validate_test_unix.go",
        "volume.go",
        "volume_copy.go",
        "volume_copy_unix.go",
        "volume_propagation_unsupported.go",
        "volume_unix.go",
        "volume_unsupported.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/docker/api/types/mount:go_default_library",
        "//vendor/github.com/docker/docker/pkg/idtools:go_default_library",
        "//vendor/github.com/docker/docker/pkg/stringid:go_default_library",
        "//vendor/github.com/opencontainers/runc/libcontainer/label:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)
