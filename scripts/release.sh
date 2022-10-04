#!/bin/bash

### Build release artifacts using Bazel.
mkdir bin \
    bin/linux/ \
    bin/linux/amd64 \
    bin/darwin \
    bin/darwin/amd64

bazelisk build \
    //:cross_linuxamd64 \
    //:cross_darwinamd64
echo

cp  bazel-out/*/bin/cross_linuxamd64_/cross_linuxamd64 bin/linux/amd64/sweep-linux-x86_64
cp  bazel-out/*/bin/cross_darwinamd64_/cross_darwinamd64 bin/darwin/amd64/sweep-darwin-x86_64

ls bin/linux/amd64 bin/darwin/amd64
