load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "download.go",
        "transfer.go",
        "upload.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/distribution:go_default_library",
        "//vendor/github.com/docker/docker/image:go_default_library",
        "//vendor/github.com/docker/docker/layer:go_default_library",
        "//vendor/github.com/docker/docker/pkg/archive:go_default_library",
        "//vendor/github.com/docker/docker/pkg/ioutils:go_default_library",
        "//vendor/github.com/docker/docker/pkg/progress:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
