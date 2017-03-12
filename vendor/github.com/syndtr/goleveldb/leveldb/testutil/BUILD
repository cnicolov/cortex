load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "db.go",
        "ginkgo.go",
        "iter.go",
        "kv.go",
        "kvtest.go",
        "storage.go",
        "util.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/ginkgo/config:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/comparer:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/errors:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/iterator:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/storage:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/util:go_default_library",
    ],
)
