FROM golang:1.10
ENV GOBIN /go/bin

RUN mkdir /app
RUN mkdir /go/src/app
ADD . /go/src/app
WORKDIR /go/src/app

# Go dep!
RUN go get -u github.com/golang/dep/...
RUN dep ensure

# Build my app
#RUN go build -o /app/main .
CMD ["go","run","main.go"]

EXPOSE 9999
