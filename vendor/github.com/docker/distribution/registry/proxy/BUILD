load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "proxyauth.go",
        "proxyblobstore.go",
        "proxymanifeststore.go",
        "proxymetrics.go",
        "proxyregistry.go",
        "proxytagservice.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution:go_default_library",
        "//vendor/github.com/docker/distribution/configuration:go_default_library",
        "//vendor/github.com/docker/distribution/context:go_default_library",
        "//vendor/github.com/docker/distribution/digest:go_default_library",
        "//vendor/github.com/docker/distribution/reference:go_default_library",
        "//vendor/github.com/docker/distribution/registry/client:go_default_library",
        "//vendor/github.com/docker/distribution/registry/client/auth:go_default_library",
        "//vendor/github.com/docker/distribution/registry/client/auth/challenge:go_default_library",
        "//vendor/github.com/docker/distribution/registry/client/transport:go_default_library",
        "//vendor/github.com/docker/distribution/registry/proxy/scheduler:go_default_library",
        "//vendor/github.com/docker/distribution/registry/storage:go_default_library",
        "//vendor/github.com/docker/distribution/registry/storage/driver:go_default_library",
    ],
)
