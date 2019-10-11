FROM golang AS builder

ADD https://github.com/golang/dep/releases/download/v0.4.1/dep-linux-amd64 /usr/bin/dep
RUN chmod +x /usr/bin/dep

#RUN go get -d github.com/golang/dep/cmd/dep && \
#    go install github.com/golang/dep/cmd/dep

WORKDIR $GOPATH/src/ipchain/

RUN CGO_ENABLED=0 go build -o /ipchain ./cli


FROM golang:alpine

COPY --from=builder /ipchain .

CMD [ "./ipchain" ]