load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["assume_role_provider.go"],
    importpath = "github.com/aws/aws-sdk-go/aws/credentials/stscreds",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/aws/aws-sdk-go/aws:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/aws/client:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/aws/credentials:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/service/sts:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["assume_role_provider_test.go"],
    importpath = "github.com/aws/aws-sdk-go/aws/credentials/stscreds",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/aws/aws-sdk-go/aws:go_default_library",
        "//vendor/github.com/aws/aws-sdk-go/service/sts:go_default_library",
        "//vendor/github.com/stretchr/testify/assert:go_default_library",
    ],
)

filegroup(
    name = "package-srcs",
    srcs = glob(["**"]),
    tags = ["automanaged"],
    visibility = ["//visibility:private"],
)

filegroup(
    name = "all-srcs",
    srcs = [":package-srcs"],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
