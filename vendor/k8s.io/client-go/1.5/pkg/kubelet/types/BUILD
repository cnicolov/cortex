load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "constants.go",
        "doc.go",
        "labels.go",
        "pod_update.go",
        "types.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/k8s.io/client-go/1.5/pkg/api:go_default_library"],
)
