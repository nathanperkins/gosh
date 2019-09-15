load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_binary(
    name = "smallsh",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

go_library(
    name = "go_default_library",
    srcs = ["smallsh.go"],
    importpath = "github.com/nathanperkins/smallsh-go/cmd",
    visibility = ["//visibility:private"],
    deps = [
        "//smallsh",
        "@github_com_sirupsen_logsrus//:go_default_library",
    ],
)
