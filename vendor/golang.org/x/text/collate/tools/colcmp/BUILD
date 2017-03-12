load("@io_bazel_rules_go//go:def.bzl", "cgo_library", "go_binary", "go_library")

cgo_library(
    name = "cgo_default_library",
    srcs = ["darwin.go"],
    clinkopts = [
        "-framework",
        "CoreFoundation",
    ],
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/golang.org/x/text/collate:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
        "//vendor/golang.org/x/text/unicode/norm:go_default_library",
    ],
)

go_library(
    name = "go_default_library",
    srcs = [
        "chars.go",
        "col.go",
        "colcmp.go",
        "gen.go",
    ],
    library = ":cgo_default_library",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/golang.org/x/text/collate:go_default_library",
        "//vendor/golang.org/x/text/language:go_default_library",
        "//vendor/golang.org/x/text/unicode/norm:go_default_library",
    ],
)

go_binary(
    name = "colcmp",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)
