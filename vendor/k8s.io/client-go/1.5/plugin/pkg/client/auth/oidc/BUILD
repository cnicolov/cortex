load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["oidc.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/coreos/go-oidc/jose:go_default_library",
        "//vendor/github.com/coreos/go-oidc/oauth2:go_default_library",
        "//vendor/github.com/coreos/go-oidc/oidc:go_default_library",
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/util/wait:go_default_library",
        "//vendor/k8s.io/client-go/1.5/rest:go_default_library",
    ],
)
