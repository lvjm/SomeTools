# 构建步骤

## 执行命令
```
docker build -t may .
```

## 执行结果
```
Sending build context to Docker daemon   5.12kB
Step 1/4 : FROM golang:1.11.5
 ---> 1454e2b3d01f
Step 2/4 : WORKDIR /go/src
 ---> Using cache
 ---> 71f5e86fa668
Step 3/4 : COPY . .
 ---> d3eab30f3516
Step 4/4 : RUN go run ./hello/main.go
 ---> Running in a0670931753a
你好，小仙女，🧚‍♀️
Removing intermediate container a0670931753a
 ---> c9a31af1b698
Successfully built c9a31af1b698
Successfully tagged maylv:latest

```
