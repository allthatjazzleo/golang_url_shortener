FROM golang:1.12

# Set go bin which doesn't appear to be set already.
# ENV GOBIN /go/bin

# build directories
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app/

# Go tidy up
RUN GO111MODULE=on go mod tidy

WORKDIR /go/src/app/cmd/app

# Build my app
RUN go build -o /app .
CMD ["/app"]