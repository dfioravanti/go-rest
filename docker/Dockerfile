############################
# STEP 1 build executable binary
############################
FROM golang:alpine AS builder

# Install git
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git
WORKDIR $GOPATH/github.com/dfioravanti/go-rest

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/app github.com/dfioravanti/go-rest/cmd/web

############################
# STEP 2 build a small image
############################
FROM scratch as prod
COPY ./db /db
COPY --from=builder /go/bin/app /go/bin/app
ENTRYPOINT ["/go/bin/app"]