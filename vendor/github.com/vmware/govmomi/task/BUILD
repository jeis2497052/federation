load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "error.go",
        "wait.go",
    ],
    importpath = "github.com/vmware/govmomi/task",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/vmware/govmomi/property:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/progress:go_default_library",
        "//vendor/github.com/vmware/govmomi/vim25/types:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["wait_test.go"],
    importpath = "github.com/vmware/govmomi/task",
    library = ":go_default_library",
    deps = ["//vendor/github.com/vmware/govmomi/vim25/types:go_default_library"],
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
