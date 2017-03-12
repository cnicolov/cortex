load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["discovery.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/emicklei/go-restful/swagger:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api/unversioned:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api/v1:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/version:go_default_library",
        "//vendor/k8s.io/client-go/1.5/testing:go_default_library",
    ],
)
