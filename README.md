# Mole
Mole is a api union for vitess and k8s

# 使用说明

1. 正常情况下可以直接go install使用，如果无法正常编译参见2

2. 如果无法正常编译需要下载对应的依赖包
``` sh
$ go get github.com/golang/dep
$ go install github.com/golang/dep/cmd/dep

# Make sure you have a go file in your directory which imports k8s.io/client-go
# first--I suggest copying one of the examples.
$ dep init
$ dep ensure k8s.io/client-go@^2.0.0

```
