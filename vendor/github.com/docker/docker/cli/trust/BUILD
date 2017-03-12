load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = ["trust.go"],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Sirupsen/logrus:go_default_library",
        "//vendor/github.com/docker/distribution/registry/client/auth:go_default_library",
        "//vendor/github.com/docker/distribution/registry/client/auth/challenge:go_default_library",
        "//vendor/github.com/docker/distribution/registry/client/transport:go_default_library",
        "//vendor/github.com/docker/docker/api/types:go_default_library",
        "//vendor/github.com/docker/docker/api/types/registry:go_default_library",
        "//vendor/github.com/docker/docker/cli/command:go_default_library",
        "//vendor/github.com/docker/docker/cliconfig:go_default_library",
        "//vendor/github.com/docker/docker/registry:go_default_library",
        "//vendor/github.com/docker/go-connections/tlsconfig:go_default_library",
        "//vendor/github.com/docker/notary:go_default_library",
        "//vendor/github.com/docker/notary/client:go_default_library",
        "//vendor/github.com/docker/notary/passphrase:go_default_library",
        "//vendor/github.com/docker/notary/storage:go_default_library",
        "//vendor/github.com/docker/notary/trustmanager:go_default_library",
        "//vendor/github.com/docker/notary/trustpinning:go_default_library",
        "//vendor/github.com/docker/notary/tuf/data:go_default_library",
        "//vendor/github.com/docker/notary/tuf/signed:go_default_library",
    ],
)
