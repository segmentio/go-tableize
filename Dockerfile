FROM golang:1.15 as builder

ARG VERSION=dev
ENV CGO_ENABLED=0

WORKDIR /build
COPY . .

RUN go build -ldflags="-X main.version=${VERSION}" .

FROM 528451384384.dkr.ecr.us-west-2.amazonaws.com/segment-scratch
COPY --from=builder /build/go-tableize  /bin/
ENTRYPOINT [ "go-tableize" ]
