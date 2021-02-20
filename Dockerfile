FROM golang:alpine AS builder

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

RUN apk add gcc g++
# Move to working directory /build
WORKDIR /build

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .
COPY ./sample.db  ./sample.db
#RUN rm -rf ./datasources/sqlite/sample_db/sample.d

# Build the application
RUN go build -o main .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp /build/main .

# Build a small image
#FROM scratch

#COPY --from=builder /dist/main /
RUN cp -r /dist/main /

# Command to run
ENTRYPOINT ["/main"]

