load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "backend_unsupported.go",
        "blobstore.go",
        "defs.go",
        "manager.go",
        "store.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/distribution/digest:go_default_library",
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/docker/docker/distribution/xfer:go_default_library",
        "//vendor/github.com/docker/docker/image:go_default_library",
        "//vendor/github.com/docker/docker/layer:go_default_library",
        "//vendor/github.com/docker/docker/libcontainerd:go_default_library",
        "//vendor/github.com/docker/docker/pkg/archive:go_default_library",
        "//vendor/github.com/docker/docker/pkg/ioutils:go_default_library",
        "//vendor/github.com/docker/docker/pkg/mount:go_default_library",
        "//vendor/github.com/docker/docker/pkg/plugingetter:go_default_library",
        "//vendor/github.com/docker/docker/pkg/plugins:go_default_library",
        "//vendor/github.com/docker/docker/pkg/progress:go_default_library",
        "//vendor/github.com/docker/docker/plugin/v2:go_default_library",
        "//vendor/github.com/docker/docker/reference:go_default_library",
        "//vendor/github.com/docker/docker/registry:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
