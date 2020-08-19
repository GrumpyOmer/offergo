# syntax = docker/dockerfile:experimental
FROM golang:1.14.2-alpine3.11 as base
FROM alpine:3.12.0 as runner

FROM base as module
ENV GOPROXY https://go.likeli.top
ENV GO111MODULE on
WORKDIR /app
ADD go.mod .
ADD go.sum .
RUN --mount=type=cache,target=/go/pkg/mod,id=offergo,sharing=locked \
	sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
	apk update && \
	apk add git && \
	go mod download -x


FROM base as builder
ENV GOPROXY https://go.likeli.top
ENV GO111MODULE on
WORKDIR /app
COPY . .
RUN --mount=type=cache,target=/go/pkg/mod,id=offergo,sharing=locked \
	sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
	apk update && \
	apk add git && \
	go build


FROM runner
WORKDIR /app
RUN  sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
	apk update && \
	apk add tzdata && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
	echo "Asia/Shanghai" > /etc/timezone && \
	apk del tzdata
EXPOSE 1234
COPY --from=builder /app/offergo /app/offergo
CMD ["./offergo"]

# syntax = docker/dockerfile:experimental