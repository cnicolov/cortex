load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "case.go",
        "delete.go",
        "expr.go",
        "insert.go",
        "part.go",
        "placeholder.go",
        "row.go",
        "select.go",
        "squirrel.go",
        "statement.go",
        "stmtcacher.go",
        "update.go",
        "where.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/lann/builder:go_default_library"],
)
