load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "events.go",
        "filter.go",
        "metrics.go",
    ],
    importpath = "github.com/docker/docker/daemon/events",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution/reference:go_default_library",
        "//vendor/github.com/docker/docker/api/types/events:go_default_library",
        "//vendor/github.com/docker/docker/api/types/filters:go_default_library",
        "//vendor/github.com/docker/docker/pkg/pubsub:go_default_library",
        "//vendor/github.com/docker/go-metrics:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["events_test.go"],
    importpath = "github.com/docker/docker/daemon/events",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/docker/api/types/events:go_default_library",
        "//vendor/github.com/docker/docker/api/types/time:go_default_library",
        "//vendor/github.com/docker/docker/daemon/events/testutils:go_default_library",
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
        "//vendor/github.com/docker/docker/daemon/events/testutils:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
