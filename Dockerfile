FROM golang:1.13-alpine AS build

ENV GO111MODULE=on
ENV GOSUMDB=off
ENV GOPROXY=direct
RUN apk add --no-cache git
WORKDIR go/src/goPhotos
ADD /* ./
RUN go mod download
RUN go build -o cmd
RUN ls -ltr

FROM scratch
COPY --from=build go/src/goPhotos/cmd /opt
ENTRYPOINT /opt/cmd