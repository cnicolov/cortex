load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "fake_authentication_client.go",
        "fake_tokenreview.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/client-go/1.5/kubernetes/typed/authentication/v1beta1:go_default_library",
        "//vendor/k8s.io/client-go/1.5/rest:go_default_library",
        "//vendor/k8s.io/client-go/1.5/testing:go_default_library",
    ],
)
