load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "endpoints.go",
        "kubernetes.go",
        "node.go",
        "pod.go",
        "service.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/github.com/prometheus/common/log:go_default_library",
        "//vendor/github.com/prometheus/common/model:go_default_library",
        "//vendor/github.com/prometheus/prometheus/config:go_default_library",
        "//vendor/github.com/prometheus/prometheus/util/strutil:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/k8s.io/client-go/1.5/kubernetes:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api/v1:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/util/runtime:go_default_library",
        "//vendor/k8s.io/client-go/1.5/rest:go_default_library",
        "//vendor/k8s.io/client-go/1.5/tools/cache:go_default_library",
    ],
)
