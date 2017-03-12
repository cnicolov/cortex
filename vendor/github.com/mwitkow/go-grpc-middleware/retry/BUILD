load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "backoff.go",
        "doc.go",
        "options.go",
        "retry.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/mwitkow/go-grpc-middleware/util/backoffutils:go_default_library",
        "//vendor/github.com/mwitkow/go-grpc-middleware/util/metautils:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
        "//vendor/golang.org/x/net/trace:go_default_library",
        "//vendor/google.golang.org/grpc:go_default_library",
        "//vendor/google.golang.org/grpc/codes:go_default_library",
        "//vendor/google.golang.org/grpc/metadata:go_default_library",
    ],
)
