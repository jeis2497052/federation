load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "auto.go",
        "regexp.go",
        "setup.go",
        "walk.go",
        "zone.go",
    ],
    importpath = "github.com/miekg/coredns/middleware/auto",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/mholt/caddy:go_default_library",
        "//vendor/github.com/miekg/coredns/core/dnsserver:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware/file:go_default_library",
        "//vendor/github.com/miekg/coredns/middleware/metrics:go_default_library",
        "//vendor/github.com/miekg/coredns/request:go_default_library",
        "//vendor/github.com/miekg/dns:go_default_library",
        "//vendor/golang.org/x/net/context:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "regexp_test.go",
        "setup_test.go",
        "walk_test.go",
        "watcher_test.go",
    ],
    importpath = "github.com/miekg/coredns/middleware/auto",
    library = ":go_default_library",
    deps = ["//vendor/github.com/mholt/caddy:go_default_library"],
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
