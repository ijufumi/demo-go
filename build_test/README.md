# Goのbuild時に変数に値を設定する方法

## 1. 値を指定する
```bash
# go build -ldflags="-X main.key=123" -o main main.go
```
をした後に、
```bash
# ./main
```
とすると、

```log
2019/09/11 23:23:06 key is 123
```
となる。

## 2. 環境変数を指定する
```bash
# go build -ldflags="-X main.key=${USER}" -o main main.go
```
をした後に、
```bash
# ./main
```
とすると、

```log
2019/09/11 23:23:06 key is ijufumi
```
となる。
