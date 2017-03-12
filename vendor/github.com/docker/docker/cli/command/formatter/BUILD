load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "container.go",
        "custom.go",
        "disk_usage.go",
        "formatter.go",
        "image.go",
        "network.go",
        "reflect.go",
        "service.go",
        "stats.go",
        "volume.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution/reference:go_default_library",
        "//vendor/github.com/docker/docker/api:go_default_library",
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/docker/docker/api/types/mount:go_default_library",
        "//vendor/github.com/docker/docker/api/types/swarm:go_default_library",
        "//vendor/github.com/docker/docker/cli/command/inspect:go_default_library",
        "//vendor/github.com/docker/docker/pkg/stringid:go_default_library",
        "//vendor/github.com/docker/docker/pkg/stringutils:go_default_library",
        "//vendor/github.com/docker/docker/reference:go_default_library",
        "//vendor/github.com/docker/docker/utils/templates:go_default_library",
        "//vendor/github.com/docker/go-units:go_default_library",
    ],
)
