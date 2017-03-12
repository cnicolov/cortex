load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["latest.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/client-go/1.5/pkg/api/unversioned:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/runtime:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/runtime/serializer/json:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/runtime/serializer/versioning:go_default_library",
        "//vendor/k8s.io/client-go/1.5/tools/clientcmd/api:go_default_library",
        "//vendor/k8s.io/client-go/1.5/tools/clientcmd/api/v1:go_default_library",
    ],
)
