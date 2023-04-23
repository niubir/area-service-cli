[English](https://github.com/niubir/area-service-cli/blob/main/helper/README-en.md) | 简体中文

## 用法

### 开始使用

1. Go

```sh
go get -u github.com/niubir/area-service-cli
```

2. 例子

```go
httpCli, _ := NewClient(HTTP, DEFALUT_HTTP_ADDRESS)
area1, _ := httpCli.GetArea()
fmt.Println(area1)

grpcCli, _ := NewClient(GRPC, DEFALUT_GRPC_ADDRESS)
area2, _ := grpcCli.GetArea()
fmt.Println(area2)
```
