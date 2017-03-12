load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "event.go",
        "force_leave.go",
        "info.go",
        "join.go",
        "keygen.go",
        "keys.go",
        "leave.go",
        "members.go",
        "monitor.go",
        "output.go",
        "query.go",
        "reachability.go",
        "rpc.go",
        "rtt.go",
        "tags.go",
        "version.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/logutils:go_default_library",
        "//vendor/github.com/hashicorp/serf/client:go_default_library",
        "//vendor/github.com/hashicorp/serf/cmd/serf/command/agent:go_default_library",
        "//vendor/github.com/hashicorp/serf/serf:go_default_library",
        "//vendor/github.com/mitchellh/cli:go_default_library",
        "//vendor/github.com/ryanuber/columnize:go_default_library",
    ],
)
