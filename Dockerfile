FROM golang:1.24
#ENV GOPROXY=https://goproxy.io,https://goproxy.cn,https://mirrors.aliyun.com/goproxy/,direct
#ARG GOPROXY=https://mirrors.aliyun.com/goproxy,direct
#ENV GOPROXY=${GOPROXY}
#ENV GOPROXY=https://mirrors.kubarcloud.com/goproxy/,direct
ENV HTTPS_PROXY=http://host.docker.internal:1080
WORKDIR /app
ENV  GOFLAGS=-x
RUN --mount=type=cache,target=/go/pkg/mod true

COPY go.mod go.sum ./
RUN --mount=type=cache,target=/go/pkg/mod \
    go mod download              # <- will now print all fetches thanks to GOFLAGS=-x
COPY . .

RUN go build -v -o server

CMD ["/app/server"]
