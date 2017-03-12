load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "fs.go",
        "image.go",
        "rootfs.go",
        "store.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/distribution/digest:go_default_library",
        "//vendor/github.com/docker/docker/api/types/container:go_default_library",
        "//vendor/github.com/docker/docker/layer:go_default_library",
        "//vendor/github.com/docker/docker/pkg/ioutils:go_default_library",
    ],
)
