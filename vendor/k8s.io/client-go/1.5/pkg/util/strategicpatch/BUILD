load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["patch.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/davecgh/go-spew/spew:go_default_library",
        "//vendor/github.com/ghodss/yaml:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/third_party/forked/golang/json:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/util/json:go_default_library",
    ],
)
