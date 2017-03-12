load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "artifactoperations.go",
        "artifactsourceoperations.go",
        "client.go",
        "costoperations.go",
        "customimageoperations.go",
        "formulaoperations.go",
        "galleryimageoperations.go",
        "laboperations.go",
        "models.go",
        "policyoperations.go",
        "policyset.go",
        "scheduleoperations.go",
        "version.go",
        "virtualmachine.go",
        "virtualnetworkoperations.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/go-autorest/autorest:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/azure:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/date:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/to:go_default_library",
    ],
)
