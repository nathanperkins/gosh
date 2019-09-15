load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
    name = "smallsh",
    srcs = [
        "cd.go",
        "smallsh.go",
    ],
    importpath = "github.com/nathanperkins/smallsh-go/smallsh",
    visibility = ["//visibility:public"],
    deps = [
        "@github_com_sirupsen_logsrus//:go_default_library",
    ],
)

go_test(
    name = "go_default_test",
    srcs = ["cd_test.go"],
    data = glob(["testdata/**"]),
    embed = [":go_default_library"],
)
