load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "doc.go",
        "event.go",
        "events_cache.go",
        "fake.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/golang/glog:go_default_library",
        "//vendor/github.com/golang/groupcache/lru:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api/errors:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api/unversioned:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/api/v1:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/runtime:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/util/clock:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/util/runtime:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/util/sets:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/util/strategicpatch:go_default_library",
        "//vendor/k8s.io/client-go/1.5/pkg/watch:go_default_library",
        "//vendor/k8s.io/client-go/1.5/rest:go_default_library",
    ],
)
