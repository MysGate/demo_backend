# demo_backend

## Compile

# Mac
```bash
GOOS=linux GOARCH=amd64 go build -o demo_backend
 ```

# Windows
```bash
go env -w  GOOS=linux
go env -w  GOARCH=amd64
go build -o demo_backe
 ```

## Go交叉编译问题(Go语言Mac/Linux/Windows下交叉编译)
https://www.cnblogs.com/xiondun/p/16971928.html


## Start 

```bash
nohup ./demo_backend -c ./config.yaml </dev/null >/dev/null 2>&1 &
```