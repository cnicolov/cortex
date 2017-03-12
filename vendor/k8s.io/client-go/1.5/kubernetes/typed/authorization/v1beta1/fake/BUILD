load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "fake_authorization_client.go",
        "fake_localsubjectaccessreview.go",
        "fake_selfsubjectaccessreview.go",
        "fake_subjectaccessreview.go",
        "fake_subjectaccessreview_expansion.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/k8s.io/client-go/1.5/kubernetes/typed/authorization/v1beta1:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/apis/authorization/v1beta1:go_default_library",
        "//vendor/k8s.io/client-go/1.5/rest:go_default_library",
        "//vendor/k8s.io/client-go/1.5/testing:go_default_library",
    ],
)
