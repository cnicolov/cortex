load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["kv.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/docker/pkg/discovery:go_default_library",
        "//vendor/github.com/docker/go-connections/tlsconfig:go_default_library",
        "//vendor/github.com/docker/libkv:go_default_library",
        "//vendor/github.com/docker/libkv/store:go_default_library",
        "//vendor/github.com/docker/libkv/store/consul:go_default_library",
        "//vendor/github.com/docker/libkv/store/etcd:go_default_library",
        "//vendor/github.com/docker/libkv/store/zookeeper:go_default_library",
    ],
)
