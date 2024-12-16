FROM alpine:latest

RUN apk add -v build-base
RUN apk add -v go 
RUN apk add -v ca-certificates
RUN apk add --no-cache \
    unzip \
    # this is needed only if you want to use scp to copy later your pb_data locally
    openssh

COPY . ./pb
WORKDIR /pb

RUN go get github.com/pocketbase/pocketbase
RUN go build
WORKDIR /

EXPOSE 8080

# start PocketBase
CMD ["/pb/pocketbase", "serve", "--http=0.0.0.0:8080"]

# FROM golang:latest
#
# WORKDIR /src
#
# # copy all your .go source files
# # (or use a .dockerignore and COPY . .)
# COPY . .
#
# COPY go.mod .
# COPY go.sum .
# RUN go mod download && go mod verify
#
# # build (switch to 1 to use the CGO SQLite)
# RUN CGO_ENABLED=0 go build -o /pocketbase
#
# # export listener port
# EXPOSE 8090
#
# # run
# # (ps. don't forget to mount the pb_data as volume in /pb_data)
# CMD ["/pocketbase", "serve", "--http=0.0.0.0:8090"]
