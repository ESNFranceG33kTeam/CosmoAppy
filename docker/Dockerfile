# Builder stage
FROM golang:1.19-alpine3.16 as builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    GOCACHE=/tmp \
    HOME=/ \
    PROJECT_NAME=CosmoAppy

LABEL org.opencontainers.image.title="${PROJECT_NAME}" \
      org.opencontainers.image.authors=${GITHUB_ACTOR} \
      org.opencontainers.image.vendor=${GITHUB_REPOSITORY} \
      org.opencontainers.image.source="https://github.com/ESNFranceG33kTeam/${PROJECT_NAME}" \
      org.opencontainers.image.url="https://github.com/ESNFranceG33kTeam/${PROJECT_NAME}/tags" \
      org.opencontainers.image.description="API for the perfect ESN association." \
      org.opencontainers.image.documentation="https://github.com/ESNFranceG33kTeam/${PROJECT_NAME}/blob/main/README.md" \
      org.opencontainers.image.os="alpine" \
      org.opencontainers.image.version=${GITHUB_REF}

RUN apk add --no-cache make gcc curl jq musl-dev

COPY . /src/
WORKDIR /src/

RUN make build-app

RUN make install-devtools
RUN make swagger

# Runner stage
FROM alpine:3.16 as runner

COPY --from=builder /src/CosmoAppy /usr/local/bin/CosmoAppy
COPY --from=builder /src/conf/swagger.yaml /usr/local/bin/conf/swagger.yaml

RUN adduser --uid 1000 --disabled-password --shell /bin/bash user
USER user

VOLUME /etc/CosmoAppy/conf/
EXPOSE 8080

WORKDIR /usr/local/bin/

ENTRYPOINT [ "/usr/local/bin/CosmoAppy" ]
CMD [ "--help" ]
