load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "docurls.go",
        "gen.go",
    ],
    visibility = ["//visibility:private"],
    deps = ["//vendor/google.golang.org/api/google-api-go-generator/internal/disco:go_default_library"],
)

go_binary(
    name = "google-api-go-generator",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
