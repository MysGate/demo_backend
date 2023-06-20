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
go build -o demo_backend
 ```


## Start 

```bash
nohup ./demo_backend -c ./config.yaml </dev/null >/dev/null 2>&1 &
```