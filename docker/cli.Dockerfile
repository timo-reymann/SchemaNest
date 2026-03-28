FROM scratch AS license
COPY LICENSE LICENSE
COPY NOTICE NOTICE

FROM busybox AS bin
COPY ./dist /dist
RUN if [[ "$(arch)" == "x86_64" ]]; then \
        architecture="amd64"; \
    else \
        architecture="arm"; \
    fi; \
    cp /dist/schema-nest-cli-linux-${architecture} /bin/schema-nest-cli && \
    chmod +x /bin/schema-nest-cli && \
    chown 1000:1000 /bin/schema-nest-cli

FROM chainguard/wolfi-base

RUN adduser -D -u 1000 schemanest
USER 1000

COPY --from=license / /

LABEL org.opencontainers.image.title="schemanest-cli" \
      org.opencontainers.image.description="Interact with the SchemaNest API with ease." \
      org.opencontainers.image.ref.name="main" \
      org.opencontainers.image.licenses='GPL-3.0' \
      org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>" \
      org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>" \
      org.opencontainers.image.url="https://github.com/timo-reymann/SchemaNest" \
      org.opencontainers.image.documentation="https://github.com/timo-reymann/SchemaNest" \
      org.opencontainers.image.source="https://github.com/timo-reymann/SchemaNest.git"

COPY --from=bin /bin/schema-nest-cli /bin/schema-nest-cli
