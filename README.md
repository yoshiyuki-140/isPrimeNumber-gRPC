# isPrimeNumber-gRPC


## Install dependencies

### Server side

```bash
# use `go mod tidy` to install dependencies
go mod tidy
```
If you want to install dependencies manually

```bash
# install grpc
go get -u google.golang.org/grpc
# install protoc's go plugin
go get -u github.com/golang/protobuf/protoc-gen-go
```

### Client side

```bash
# if you want to activate virtual environment
python -m venv venv
# install dependencies
python -m pip -r requirements.txt
```

## Usage

### Server side
```bash
cd server
go run server.go
```

### Client side
```bash
cd client
python client.py
```

## 参考文献

- messageに使える型
  - https://qiita.com/yukina-ge/items/98fe190cef2024d45fbd

- ハンズオン-go
  - https://qiita.com/ishishow/items/c9b633d3d4d22db74a6d

- ハンズオン-py
  - https://zenn.dev/sugasuga/articles/10e9843263f215
