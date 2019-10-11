FROM golang AS builder

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

WORKDIR $GOPATH/src/ipchain/
COPY Gopkg.toml Gopkg.lock ./
RUN dep ensure --vendor-only
COPY . ./
RUN CGO_ENABLED=0 go build -o /ipchain ./cli

FROM golang:alpine

COPY --from=builder /ipchain .

CMD [ "./ipchain" ]