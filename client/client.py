import grpc
from isPrime import isPrime_pb2, isPrime_pb2_grpc
import os

# os.environ["GRPC_VERBOSITY"] = "DEBUG"
os.environ["GRPC_VERBOSITY"] = "NONE"  # ログを非表示にする
os.environ["NO_PROXY"] = "localhost"


def main():
    """メイン関数"""
    # サーバーのアドレスとポートを指定
    with grpc.insecure_channel("localhost:9000") as channel:
        stub = isPrime_pb2_grpc.IsPrimeFuncStub(channel)
        # サーバーに送るメッセージを定義
        response = stub.CheckPrime(isPrime_pb2.Value(Value=askNumber()))
        print("Is 17 a prime number? " + ("Yes" if response.IsPrime else "No"))


def askNumber():
    """数字を聞いて返すだけの関数

    Returns:
        int: 素数かどうか調べたい数字
    """
    num = input("Please input a number: ")
    return int(num)


if __name__ == "__main__":
    main()
