load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "go_default_library",
    srcs = [
        "analysis.go",
        "api.go",
        "callgraph.go",
        "constraint.go",
        "doc.go",
        "gen.go",
        "hvn.go",
        "intrinsics.go",
        "labels.go",
        "opt.go",
        "print.go",
        "query.go",
        "reflect.go",
        "solve.go",
        "util.go",
    ],
    importpath = "golang.org/x/tools/go/pointer",
    visibility = ["//visibility:public"],
    deps = [
        "//vendor/golang.org/x/tools/container/intsets:go_default_library",
        "//vendor/golang.org/x/tools/go/callgraph:go_default_library",
        "//vendor/golang.org/x/tools/go/ssa:go_default_library",
        "//vendor/golang.org/x/tools/go/types/typeutil:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = [
        "query_test.go",
        "stdlib_test.go",
    ],
    importpath = "golang.org/x/tools/go/pointer",
    library = ":go_default_library",
    deps = [
        "//vendor/golang.org/x/tools/go/buildutil:go_default_library",
        "//vendor/golang.org/x/tools/go/loader:go_default_library",
        "//vendor/golang.org/x/tools/go/ssa:go_default_library",
        "//vendor/golang.org/x/tools/go/ssa/ssautil:go_default_library",
    ],
)

go_test(
    name = "go_default_xtest",
    srcs = [
        "example_test.go",
        "pointer_test.go",
    ],
    importpath = "golang.org/x/tools/go/pointer_test",
    deps = [
        ":go_default_library",
        "//vendor/golang.org/x/tools/go/callgraph:go_default_library",
        "//vendor/golang.org/x/tools/go/loader:go_default_library",
        "//vendor/golang.org/x/tools/go/ssa:go_default_library",
        "//vendor/golang.org/x/tools/go/ssa/ssautil:go_default_library",
        "//vendor/golang.org/x/tools/go/types/typeutil:go_default_library",
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
        "//vendor/golang.org/x/tools/go/pointer/testdata:all-srcs",
    ],
    tags = ["automanaged"],
    visibility = ["//visibility:public"],
)
