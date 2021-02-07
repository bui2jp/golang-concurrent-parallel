# golang x 平行(concurrent) x 並列(parallel)

golangの平行・並列処理のサンプル

##　Docker(golang:1.14.13)を利用

```
% docker run -it --rm golang:1.14.13 go version
```

## 1.平行 (concurrent)
```
% docker run -it --rm -v $PWD:/root -w /root golang:1.14.13 go run 01_concurrent/main.go
```

## 2.並列 (parallel)
```
% docker run -it --rm -v $PWD:/root -w /root golang:1.14.13 go run 02_parallel/main.go
```
