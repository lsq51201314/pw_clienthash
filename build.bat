@echo off
cd %cd%
echo 正在编译Windows平台 。。。
go build -ldflags "-s -w" main.go
echo 正在编译Linux平台 。。。
SET CGO_ENABLED=0
SET GOARCH=amd64
SET GOOS=linux
go build -ldflags "-s -w" main.go
echo 编译完成 。。。
pause