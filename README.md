# isprimeNumber-gRPC-go


## Install dependencies

### Server side

```bash
# install grpc
go get -u google.golang.org/grpc
# install protoc's go plugin
go get -u github.com/golang/protobuf/protoc-gen-go
```

### Client side

```bash
python -m pip venv venv
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
