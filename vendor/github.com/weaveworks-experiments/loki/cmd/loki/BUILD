load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/prometheus/prometheus/config:go_default_library",
        "//vendor/github.com/prometheus/prometheus/retrieval:go_default_library",
        "//vendor/github.com/weaveworks-experiments/loki/pkg/api:go_default_library",
        "//vendor/github.com/weaveworks-experiments/loki/pkg/scraper:go_default_library",
        "//vendor/github.com/weaveworks-experiments/loki/pkg/storage:go_default_library",
        "//vendor/github.com/weaveworks-experiments/loki/pkg/zipkin-ui:go_default_library",
        "//vendor/github.com/weaveworks/common/logging:go_default_library",
        "//vendor/github.com/weaveworks/common/server:go_default_library",
    ],
)

go_binary(
    name = "loki",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
