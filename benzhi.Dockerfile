FROM golang:1.22-bookworm
WORKDIR /workspace
ENV GOPROXY=https://goproxy.cn,direct
ENV GOTOOLCHAIN=local
COPY . .
RUN go mod download
RUN go build -o /usr/local/bin/mireflux .
ENTRYPOINT ["/usr/local/bin/mireflux"]
