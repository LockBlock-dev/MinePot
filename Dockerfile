ARG GOLANG_VERSION=1.20

FROM golang:${GOLANG_VERSION}-alpine as build

WORKDIR /app

COPY ./handler handler/
COPY ./types types/
COPY ./util util/

RUN set -eux; \
    go mod download; \
    go build -ldflags "-s -w" -o /bin/minepot;



FROM alpine

ENV PORT=25565

WORKDIR /app

COPY ./config.json .
COPY ./assets assets/
COPY --from=build /bin/minepot /bin/minepot

RUN set -eux; \
    # Create the MinePot directory
    mkdir -p /etc/minepot; \
    # Copy the config
    cp config.json /etc/minepot/

EXPOSE ${PORT}

CMD [ "/bin/minepot" ]
