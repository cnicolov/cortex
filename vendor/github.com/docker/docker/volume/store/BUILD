load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "db.go",
        "errors.go",
        "restore.go",
        "store.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/boltdb/bolt:go_default_library",
        "//vendor/github.com/docker/docker/pkg/locker:go_default_library",
        "//vendor/github.com/docker/docker/volume:go_default_library",
        "//vendor/github.com/docker/docker/volume/drivers:go_default_library",
        "//vendor/github.com/pkg/errors:go_default_library",
    ],
)
