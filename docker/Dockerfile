# Builder stage
FROM golang:1.19-alpine3.16 as builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    GOCACHE=/tmp \
    HOME=/go \
    PROJECT_NAME=sAPI

LABEL org.opencontainers.image.title="${PROJECT_NAME}}" \
      org.opencontainers.image.authors=${GITHUB_ACTOR} \
      org.opencontainers.image.vendor=${GITHUB_REPOSITORY} \
      org.opencontainers.image.source="https://github.com/ESNFranceG33kTeam/${PROJECT_NAME}" \
      org.opencontainers.image.url="https://github.com/ESNFranceG33kTeam/${PROJECT_NAME}/tags" \
      org.opencontainers.image.description="API for the perfect ESN association." \
      org.opencontainers.image.documentation="https://github.com/ESNFranceG33kTeam/${PROJECT_NAME}/blob/main/README.md" \
      org.opencontainers.image.os="alpine" \
      org.opencontainers.image.version=${GITHUB_REF}

RUN apk add --no-cache make gcc curl jq

COPY . /src/
WORKDIR /src/

RUN make build-app

RUN download_url=$(curl -s https://api.github.com/repos/go-swagger/go-swagger/releases/latest | \
    jq -r '.assets[] | select(.name | contains("'"$(uname | tr '[:upper:]' '[:lower:]')"'_amd64")) | .browser_download_url') \
    && curl -o /usr/local/bin/swagger -L'#' "$download_url" \
    && chmod +x /usr/local/bin/swagger
RUN make swagger

# Runner stage
FROM alpine:3.16 as runner

COPY --from=builder /src/bin/${PROJECT_NAME} /usr/local/bin/${PROJECT_NAME}
COPY --from=builder /src/swagger.yaml /usr/local/bin/swagger.yaml

RUN adduser --uid 1000 --disabled-password --shell /bin/bash user
USER user

VOLUME /etc/${PROJECT_NAME}/conf/
EXPOSE 8080

WORKDIR /usr/local/bin/

ENTRYPOINT [ "/usr/local/bin/sAPI" ]
CMD [ "--help" ]
