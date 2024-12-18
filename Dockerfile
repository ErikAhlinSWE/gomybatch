FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .

RUN go get -d -v

RUN go build -o /app/cmd/batch
FROM scratch

COPY --from=builder /app/cmd/batch /batch
COPY *.yml ./ 


ENTRYPOINT ["/batch"]