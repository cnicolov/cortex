load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "backdoor.go",
        "gccgoinstallation.go",
        "importer.go",
        "parser.go",
    ],
    visibility = ["//vendor/golang.org/x/tools/go:__subpackages__"],
)
