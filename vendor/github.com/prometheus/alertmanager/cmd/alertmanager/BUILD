load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/prometheus/alertmanager/api:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/config:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/dispatch:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/inhibit:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/nflog:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/notify:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/provider/mem:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/silence:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/template:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/types:go_default_library",
        "//vendor/github.com/prometheus/alertmanager/ui:go_default_library",
        "//vendor/github.com/prometheus/client_golang/prometheus:go_default_library",
        "//vendor/github.com/prometheus/common/log:go_default_library",
        "//vendor/github.com/prometheus/common/route:go_default_library",
        "//vendor/github.com/prometheus/common/version:go_default_library",
        "//vendor/github.com/weaveworks/mesh:go_default_library",
    ],
)

go_binary(
    name = "alertmanager",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
