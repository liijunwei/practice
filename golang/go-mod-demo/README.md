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

+ 为当前module添加一个依赖
    + import
    + go mod tidy





