load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["valuecollector.go"],
    visibility = ["//vendor/cloud.google.com/go/cmd/go-cloud-debug-agent:__subpackages__"],
    deps = [
        "//vendor/golang.org/x/debug:go_default_library",
        "//vendor/google.golang.org/api/clouddebugger/v2:go_default_library",
    ],
)
