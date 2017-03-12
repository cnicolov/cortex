load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "certificateorders.go",
        "certificates.go",
        "classicmobileservices.go",
        "client.go",
        "domains.go",
        "global.go",
        "globalcertificateorder.go",
        "globaldomainregistration.go",
        "globalresourcegroups.go",
        "hostingenvironments.go",
        "managedhostingenvironments.go",
        "models.go",
        "provider.go",
        "recommendations.go",
        "serverfarms.go",
        "sites.go",
        "topleveldomains.go",
        "usageoperations.go",
        "version.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/Azure/go-autorest/autorest:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/azure:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/date:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/to:go_default_library",
        "//vendor/github.com/Azure/go-autorest/autorest/validation:go_default_library",
    ],
)
