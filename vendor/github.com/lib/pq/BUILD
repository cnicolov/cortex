load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "array.go",
        "buf.go",
        "conn.go",
        "conn_go18.go",
        "copy.go",
        "doc.go",
        "encode.go",
        "error.go",
        "notify.go",
        "ssl.go",
        "ssl_go1.7.go",
        "ssl_permissions.go",
        "url.go",
        "user_posix.go",
        "uuid.go",
    ],
    visibility = ["//visibility:public"],
    deps = ["//vendor/github.com/lib/pq/oid:go_default_library"],
)
