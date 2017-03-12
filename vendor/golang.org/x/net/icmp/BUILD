load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "dstunreach.go",
        "echo.go",
        "endpoint.go",
        "extension.go",
        "helper.go",
        "helper_posix.go",
        "interface.go",
        "ipv4.go",
        "ipv6.go",
        "listen_posix.go",
        "message.go",
        "messagebody.go",
        "mpls.go",
        "multipart.go",
        "packettoobig.go",
        "paramprob.go",
        "timeexceeded.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/net/internal/iana:go_default_library",
        "//vendor/golang.org/x/net/ipv4:go_default_library",
        "//vendor/golang.org/x/net/ipv6:go_default_library",
    ],
)
