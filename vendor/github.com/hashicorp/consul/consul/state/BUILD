load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "acl.go",
        "catalog.go",
        "coordinate.go",
        "delay.go",
        "graveyard.go",
        "kvs.go",
        "notify.go",
        "prepared_query.go",
        "prepared_query_index.go",
        "schema.go",
        "session.go",
        "state_store.go",
        "tombstone_gc.go",
        "txn.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/hashicorp/consul/consul/prepared_query:go_default_library",
        "//vendor/github.com/hashicorp/consul/consul/structs:go_default_library",
        "//vendor/github.com/hashicorp/consul/types:go_default_library",
        "//vendor/github.com/hashicorp/go-memdb:go_default_library",
        "//vendor/github.com/hashicorp/serf/coordinate:go_default_library",
    ],
)
