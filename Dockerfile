############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"'  -o /go/bin/hello

FROM scratch
COPY --from=builder /go/bin/hello /go/bin/hello
ENTRYPOINT ["/go/bin/hello"]
