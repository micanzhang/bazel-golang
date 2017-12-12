# bazel-golang[![Build Status](https://travis-ci.org/micanzhang/bazel-golang.svg?branch=master)](https://travis-ci.org/micanzhang/bazel-golang)
docker image include bazel and golang binaray

+ bazel: [0.5.4](https://github.com/bazelbuild/bazel/releases/tag/0.5.4)
+ golang: [1.9.2](https://golang.org/dl/)

## Install

### Build

build docker image:

```sh
bazel build -t micanzhang/bazel-golang -f Dockerfile .
```

### Pull

pull image from [docker hub](https://hub.docker.com/r/micanzhang/bazel-golang/):

```sh
docker pull micanzhang/bazel-golang
```

### How write bazel build script?

please visit: [https://github.com/lingochamp/rules_go](https://github.com/lingochamp/rules_go)


