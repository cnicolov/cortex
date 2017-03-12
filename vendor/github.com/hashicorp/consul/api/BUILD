load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "acl.go",
        "agent.go",
        "api.go",
        "catalog.go",
        "coordinate.go",
        "event.go",
        "health.go",
        "kv.go",
        "lock.go",
        "operator.go",
        "prepared_query.go",
        "raw.go",
        "semaphore.go",
        "session.go",
        "snapshot.go",
        "status.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/go-cleanhttp:go_default_library",
        "//vendor/github.com/hashicorp/serf/coordinate:go_default_library",
    ],
)
