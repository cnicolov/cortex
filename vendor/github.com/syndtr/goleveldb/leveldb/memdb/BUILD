load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["memdb.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/syndtr/goleveldb/leveldb/comparer:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/errors:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/iterator:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/util:go_default_library",
    ],
)
