load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/k8s.io/client-go/1.5/kubernetes:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api:go_default_library",
        "//vendor/k8s.io/client-go/1.5/tools/clientcmd:go_default_library",
    ],
)

go_binary(
    name = "out-of-cluster",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
