# compile env
FROM golang:1.11
COPY . /amm-bots
WORKDIR /amm-bots

# compile main.go
RUN go build -o bin/amm-bots -v -ldflags '-s -w' main.go

# execute env
FROM alpine
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN apk --no-cache add ca-certificates

# copy binary file from compile env to execute env
COPY --from=0 /amm-bots/bin/amm-bots /bin/
ENTRYPOINT /bin/amm-bots
