FROM golang:1.13-alpine AS build

WORKDIR /src/
COPY . /src/

RUN CGO_ENABLED=0 go build -o /bin/cmd

FROM scratch
COPY --from=build /bin/cmd /bin/cmd
ENTRYPOINT ["/bin/cmd"]