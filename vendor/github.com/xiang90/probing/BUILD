load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "prober.go",
        "server.go",
        "status.go",
    ],
    importpath = "github.com/xiang90/probing",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_test",
    srcs = ["prober_test.go"],
    importpath = "github.com/xiang90/probing",
    library = ":go_default_library",
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
