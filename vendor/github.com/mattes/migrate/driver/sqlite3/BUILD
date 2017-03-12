load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["sqlite3.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/mattes/migrate/driver:go_default_library",
        "//vendor/github.com/mattes/migrate/file:go_default_library",
        "//vendor/github.com/mattes/migrate/migrate/direction:go_default_library",
        "//vendor/github.com/mattn/go-sqlite3:go_default_library",
    ],
)
