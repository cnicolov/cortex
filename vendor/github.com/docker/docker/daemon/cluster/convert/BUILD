load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "container.go",
        "network.go",
        "node.go",
        "secret.go",
        "service.go",
        "swarm.go",
        "task.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/docker/docker/api/types/container:go_default_library",
        "//vendor/github.com/docker/docker/api/types/mount:go_default_library",
        "//vendor/github.com/docker/docker/api/types/network:go_default_library",
        "//vendor/github.com/docker/docker/api/types/swarm:go_default_library",
        "//vendor/github.com/docker/docker/pkg/namesgenerator:go_default_library",
        "//vendor/github.com/docker/swarmkit/api:go_default_library",
        "//vendor/github.com/docker/swarmkit/protobuf/ptypes:go_default_library",
    ],
)
