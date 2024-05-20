# syntax=docker/dockerfile:1

FROM golang:1.22.0-alpine as base
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux
# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY . /goapps/

WORKDIR /goapps

FROM base AS build-logservice
RUN go build -o /bin/logservice ./cmd/logservice

FROM scratch AS Prod
COPY --from=build-logservice /bin/logservice /bin/
# Run
ENTRYPOINT ["/bin/logservice"]