load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = ["redis.go"],
    importpath = "github.com/docker/distribution/registry/storage/cache/redis",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/github.com/docker/distribution:go_default_library",
        "//vendor/github.com/docker/distribution/context:go_default_library",
        "//vendor/github.com/docker/distribution/reference:go_default_library",
        "//vendor/github.com/docker/distribution/registry/storage/cache:go_default_library",
        "//vendor/github.com/garyburd/redigo/redis:go_default_library",
        "//vendor/github.com/opencontainers/go-digest:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["redis_test.go"],
    importpath = "github.com/docker/distribution/registry/storage/cache/redis",
    library = ":go_default_library",
    deps = [
        "//vendor/github.com/docker/distribution/registry/storage/cache/cachecheck:go_default_library",
        "//vendor/github.com/garyburd/redigo/redis:go_default_library",
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
