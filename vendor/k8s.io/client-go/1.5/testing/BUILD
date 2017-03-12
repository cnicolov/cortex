load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "actions.go",
        "fake.go",
        "fixture.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/client-go/1.5/pkg/api:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api/errors:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api/meta:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api/unversioned:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/apimachinery/registered:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/fields:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/labels:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/runtime:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/version:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/watch:go_default_library",
        "//vendor/k8s.io/client-go/1.5/rest:go_default_library",
    ],
)
