load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "portforward.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/client-go/1.5/pkg/api:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/kubelet/server/portforward:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/util/httpstream:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/util/runtime:go_default_library",
    ],
)
