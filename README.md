# minio-sdk-gomobile
minio go mobile sdk

```shell

go get github.com/go-resty/resty/v2
go get golang.org/x/mobile/cmd/gomobile
go get golang.org/x/mobile/bind

gomobile init

```

```shell
#go get -u golang.org/x/mobile/cmd/gomobile
gomobile bind -target ios ./cgominio
rm -rf Cgominio.xcframework.tar.gz 
tar -zcvf Cgominio.xcframework.tar.gz Cgominio.xcframework
```