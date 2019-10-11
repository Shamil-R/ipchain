FROM golang AS builder

RUN go get -d github.com/golang/dep/cmd/dep && \
    go install github.com/golang/dep/cmd/dep

WORKDIR $GOPATH/src/github.com/Shamil-R/ipchain

COPY Gopkg.toml Gopkg.lock ./

RUN dep ensure --vendor-only -v

COPY . ./

RUN CGO_ENABLED=0 go build -o /ipchain ./


FROM golang:alpine

COPY --from=builder /ipchain .
COPY --from=builder $GOPATH/src/github.com/Shamil-R/ipchain/config ./config

CMD [ "./ipchain", "qa" ]