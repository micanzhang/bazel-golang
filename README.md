# bazel-golang
docker image include bazel and golang binaray

+ bazel: 0.5.4
+ golang: 1.9.2

## Install

### Build

build docker image:

```sh
bazel build -t micanzhang/bazel-golang -f Dockerfile .
```

### Pull

pull image from docker hub:

```sh
docker pull micanzhang/bazel-golang
```

### How write bazel build script?

please visit: [https://github.com/lingochamp/rules_go](https://github.com/lingochamp/rules_go)



