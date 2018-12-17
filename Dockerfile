# build stage
FROM golang:alpine AS build-env
LABEL stage=intermediate
COPY . /go/src/github.com/sjljrvis/gArch/
WORKDIR /go/src/github.com/sjljrvis/gArch

RUN go build -o gArch
ENV PORT 3000
EXPOSE 3000


# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env go/src/github.com/sjljrvis/gArch/gArch /app/
ENTRYPOINT ./gArch
