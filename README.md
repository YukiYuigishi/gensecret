# gensecret

任意のバイト数のシークレットを生成するコマンド

デフォルトのエンコードは `hex` です（`-enc`で指定可能）。

## Usage

```
gensecret -n 32 -enc hex
gensecret -n 32 -enc base64
gensecret -n 32 -enc base64url
```
