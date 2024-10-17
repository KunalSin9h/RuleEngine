FROM  golang:alpine AS bin_build

WORKDIR /ruleengine

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY cmd ./cmd
COPY internal ./internal

RUN go build -o bin  ./cmd/api/*.go

FROM alpine:latest

WORKDIR /ruleengine
COPY ui ./ui
COPY schema ./schema
COPY --from=bin_build /ruleengine/bin .

ENTRYPOINT ["/ruleengine/bin"]
