import grpc
from isPrime import isPrime_pb2, isPrime_pb2_grpc
import os

# os.environ["GRPC_VERBOSITY"] = "DEBUG"
os.environ["GRPC_VERBOSITY"] = "NONE"  # ログを非表示にする
os.environ["NO_PROXY"] = "localhost"


def run():
    # サーバーのアドレスとポートを指定
    with grpc.insecure_channel("localhost:9000") as channel:
        stub = isPrime_pb2_grpc.IsPrimeFuncStub(channel)
        # サーバーに送るメッセージを定義
        response = stub.CheckPrime(isPrime_pb2.Value(Value=17))
        print("Is 17 a prime number? " + ("Yes" if response.IsPrime else "No"))


if __name__ == "__main__":
    run()
