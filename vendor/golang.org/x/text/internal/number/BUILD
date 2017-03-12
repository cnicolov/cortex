load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "common.go",
        "decimal.go",
        "extfloat.go",
        "ftoa.go",
        "itoa.go",
        "number.go",
        "pattern.go",
        "plural.go",
        "tables.go",
    ],
    visibility = ["//vendor/golang.org/x/text:__subpackages__"],
    deps = [
        "//vendor/golang.org/x/text/internal:go_default_library",
        "//vendor/golang.org/x/text/internal/format/plural:go_default_library",
        "//vendor/golang.org/x/text/internal/stringset:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
    ],
)
