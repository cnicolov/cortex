load("@io_bazel_rules_go//go:def.bzl", "go_library")

go_library(
    name = "go_default_library",
    srcs = [
        "t.go",
        "thetest.pb.go",
        "uuid.go",
    ],
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/gogo/protobuf/gogoproto:go_default_library",
        "//vendor/github.com/gogo/protobuf/proto:go_default_library",
        "//vendor/github.com/gogo/protobuf/protoc-gen-gogo/descriptor:go_default_library",
        "//vendor/github.com/gogo/protobuf/sortkeys:go_default_library",
        "//vendor/github.com/gogo/protobuf/test/custom:go_default_library",
        "//vendor/github.com/gogo/protobuf/test/custom-dash-type:go_default_library",
    ],
)

filegroup(
    name = "go_default_library_protos",
    srcs = ["thetest.proto"],
    visibility = ["//visibility:public"],
)
