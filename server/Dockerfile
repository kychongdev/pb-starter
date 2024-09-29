FROM golang:latest

WORKDIR /src

# copy all your .go source files
# (or use a .dockerignore and COPY . .)
COPY . .

COPY go.mod .
COPY go.sum .
RUN go mod download && go mod verify

# build (switch to 1 to use the CGO SQLite)
RUN CGO_ENABLED=0 go build -o /pocketbase

# export listener port
EXPOSE 8090

# run
# (ps. don't forget to mount the pb_data as volume in /pb_data)
CMD ["/pocketbase", "serve", "--http=0.0.0.0:8090"]
