load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["bidirule.go"],
    importpath = "golang.org/x/text/secure/bidirule",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/text/transform:go_default_library",
        "//vendor/golang.org/x/text/unicode/bidi:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "bench_test.go",
        "bidirule_test.go",
    ],
    importpath = "golang.org/x/text/secure/bidirule",
    library = ":go_default_library",
    deps = [
        "//vendor/golang.org/x/text/internal/testtext:go_default_library",
        "//vendor/golang.org/x/text/transform:go_default_library",
        "//vendor/golang.org/x/text/unicode/bidi:go_default_library",
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
