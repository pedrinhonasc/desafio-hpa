FROM golang:alpine as builder

WORKDIR /go/src

COPY ./main.go .

RUN CGO_ENABLED=0 GOOS=linux go build -o /go/bin/sqrt-loop main.go -ldflags="-s -w"

FROM scratch

COPY --from=builder /go/bin/sqrt-loop /go/bin/sqrt-loop

ENTRYPOINT [ "/go/bin/sqrt-loop"]

EXPOSE 8000