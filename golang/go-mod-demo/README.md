# bootstrap

```bash
mkdir go-mod-demo
cd go-mod-demo
touch main.go
edit main.go
go mod init github.com/liijunwei/go-mod-demo
go mod tidy

go build # error
go get -u golang.org/x/sys # https://stackoverflow.com/questions/71507321/go-1-18-build-error-on-mac-unix-syscall-darwin-1-13-go253-golinkname-mus

./go-mod-demo
```

# go module 的6类常规操作

1. 为当前module添加一个依赖
```bash
# import package in go file
go mod tidy
# or
go get package-name
```

2. 升级/降级依赖的版本
```bash
go list -m -versions github.com/sirupsen/logrus
go get github.com/sirupsen/logrus@v1.7.1

# or

go mod edit -require=github.com/sirupsen/logrus@v1.7.1
go mod tidy
```

3. 添加一个主版本号大于1的依赖(unclear)

4. 升级版本到一个不兼容的版本

5. 移除一个依赖
```bash
go list -m all | grep redis
# remove go-redis in main.go
go build
go list -m all | grep redis # go-redis 还在

go mod tidy
go list -m all | grep redis # go-redis 被清理掉了
```

6. 特殊情况: 使用vendor

