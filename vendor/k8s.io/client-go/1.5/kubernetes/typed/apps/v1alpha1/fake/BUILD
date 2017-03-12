load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "fake_apps_client.go",
        "fake_petset.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/client-go/1.5/kubernetes/typed/apps/v1alpha1:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api/unversioned:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/apis/apps/v1alpha1:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/labels:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/watch:go_default_library",
        "//vendor/k8s.io/client-go/1.5/rest:go_default_library",
        "//vendor/k8s.io/client-go/1.5/testing:go_default_library",
    ],
)
