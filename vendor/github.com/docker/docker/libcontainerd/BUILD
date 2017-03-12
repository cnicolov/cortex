load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "client.go",
        "container.go",
        "pausemonitor_unix.go",
        "process.go",
        "remote.go",
        "types.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/containerd/api/grpc/types:go_default_library",
        "//vendor/github.com/docker/docker/pkg/locker:go_default_library",
        "//vendor/github.com/opencontainers/runtime-spec/specs-go:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)
