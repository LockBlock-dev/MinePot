ARG GOLANG_VERSION=1.20

FROM golang:${GOLANG_VERSION}-alpine

ENV PORT=25565

WORKDIR /app

COPY ./src src/
COPY ./config.json .
COPY ./assets assets/

RUN set -eux; \
    cd src; \
    go mod download; \
    go build -ldflags "-s -w" -o /minepot; \
    cd ..

RUN set -eux; \
    # Create the MinePot directory
    mkdir -p /etc/minepot; \
    # Copy the config
    cp config.json /etc/minepot/

EXPOSE ${PORT}

CMD [ "/minepot" ]
