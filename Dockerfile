FROM golang:1.22-alpine AS build

WORKDIR /src
COPY go.mod ./
COPY cmd ./cmd

ARG APP_VERSION=dev
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build \
    -ldflags "-s -w" \
    -o /out/server ./cmd/server

FROM gcr.io/distroless/static-debian12:nonroot

ENV PORT=8080
ARG APP_VERSION=dev
ENV APP_VERSION=${APP_VERSION}
EXPOSE 8080

COPY --from=build /out/server /server
USER nonroot:nonroot
ENTRYPOINT ["/server"]
