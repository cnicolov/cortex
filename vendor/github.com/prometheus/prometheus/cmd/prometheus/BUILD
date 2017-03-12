load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "config.go",
        "main.go",
    ],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/asaskevich/govalidator:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/github.com/prometheus/common/log:go_default_library",
        "//vendor/github.com/prometheus/common/model:go_default_library",
        "//vendor/github.com/prometheus/common/version:go_default_library",
        "//vendor/github.com/prometheus/prometheus/config:go_default_library",
        "//vendor/github.com/prometheus/prometheus/notifier:go_default_library",
        "//vendor/github.com/prometheus/prometheus/promql:go_default_library",
        "//vendor/github.com/prometheus/prometheus/retrieval:go_default_library",
        "//vendor/github.com/prometheus/prometheus/rules:go_default_library",
        "//vendor/github.com/prometheus/prometheus/storage:go_default_library",
        "//vendor/github.com/prometheus/prometheus/storage/local:go_default_library",
        "//vendor/github.com/prometheus/prometheus/storage/local/chunk:go_default_library",
        "//vendor/github.com/prometheus/prometheus/storage/local/index:go_default_library",
        "//vendor/github.com/prometheus/prometheus/storage/remote:go_default_library",
        "//vendor/github.com/prometheus/prometheus/web:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

go_binary(
    name = "prometheus",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
