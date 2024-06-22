# isPrimeNumber-gRPC

gRPCの実験として、素数判定のプロシージャを呼ぶサンプルを作成しました。
server-pyはPythonのバージョンのgRPCサーバーで、server-goはGoのバージョンのgRPCサーバーです。
client-pyはPythonのバージョンのgRPCクライアントです。
もし、Pythonのみで実行したい場合は、server-goを実行する必要はありませんし、
Goの環境構築を行う必要はありません。<br>
以下はセットアップの方法です。

### Server side Setup (Go)

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
# 
```


### Server side Setup (Python)

```bash
# if you want to activate virtual environment
python -m venv venv
# install dependencies
python -m pip -r requirements.txt
cd server-py
# execute
python server.py
```

### Client side Setup (Python)

```bash
# if you want to activate virtual environment
python -m venv venv
# install dependencies
python -m pip -r requirements.txt
cd client-py
# execute
python client.py
```

## Usage

### Server side
```bash
cd server-go
go run server.go
```

Or

```bash
cd server-py
python server.py
```

### Client side

```bash
cd client-py
python client.py
```

## 参考文献

- messageに使える型
  - https://qiita.com/yukina-ge/items/98fe190cef2024d45fbd

- ハンズオン-go
  - https://qiita.com/ishishow/items/c9b633d3d4d22db74a6d

- ハンズオン-py
  - https://zenn.dev/sugasuga/articles/10e9843263f215
