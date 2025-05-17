@echo off
REM 编译为 Linux 64位可执行文件
set GOOS=linux
set GOARCH=amd64
set CGO_ENABLED=0

echo 正在编译 Linux 版可执行文件...
go build -o match main.go

if errorlevel 1 (
    echo 编译失败!
    exit /b 1
)

echo 编译成功，生成文件: match