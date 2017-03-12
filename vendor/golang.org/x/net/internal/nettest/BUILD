load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "helper_bsd.go",
        "helper_posix.go",
        "helper_unix.go",
        "interface.go",
        "rlimit.go",
        "stack.go",
    ],
    visibility = ["//vendor/golang.org/x/net:__subpackages__"],
)
