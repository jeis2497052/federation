load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/containernetworking/cni/plugins/test/noop",
    visibility = ["//visibility:private"],
    deps = [
        "//vendor/github.com/containernetworking/cni/pkg/skel:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/types/current:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/version:go_default_library",
        "//vendor/github.com/containernetworking/cni/plugins/test/noop/debug:go_default_library",
    ],
)

go_binary(
    name = "noop",
    importpath = "github.com/containernetworking/cni/plugins/test/noop",
    library = ":go_default_library",
    visibility = ["//visibility:public"],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "noop_suite_test.go",
        "noop_test.go",
    ],
    importpath = "github.com/containernetworking/cni/plugins/test/noop_test",
    deps = [
        "//vendor/github.com/containernetworking/cni/pkg/skel:go_default_library",
        "//vendor/github.com/containernetworking/cni/pkg/version:go_default_library",
        "//vendor/github.com/containernetworking/cni/plugins/test/noop/debug:go_default_library",
        "//vendor/github.com/onsi/ginkgo:go_default_library",
        "//vendor/github.com/onsi/gomega:go_default_library",
        "//vendor/github.com/onsi/gomega/gexec:go_default_library",
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
    srcs = [
        ":package-srcs",
        "//vendor/github.com/containernetworking/cni/plugins/test/noop/debug:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
