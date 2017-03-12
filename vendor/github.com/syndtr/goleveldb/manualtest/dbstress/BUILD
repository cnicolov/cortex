load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "key.go",
        "main.go",
    ],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/syndtr/goleveldb/leveldb:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/errors:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/opt:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/storage:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/table:go_default_library",
        "//vendor/github.com/syndtr/goleveldb/leveldb/util:go_default_library",
    ],
)

go_binary(
    name = "dbstress",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
