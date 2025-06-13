FROM busybox AS bin
COPY ./dist /dist
RUN if [[ "$(arch)" == "x86_64" ]]; then \
        architecture="amd64"; \
    else \
        architecture="arm"; \
    fi; \
    cp /dist/schema-nest-cli-linux-${architecture} /bin/schemanest-cli && \
    chmod +x /bin/schemanest-cli && \
    chown 1000:1000 /bin/schemanest-cli

FROM chainguard/wolfi-base
LABEL org.opencontainers.image.title="schemanest-cli"
LABEL org.opencontainers.image.description="Interact with the SchemaNest API with ease."
LABEL org.opencontainers.image.ref.name="main"
LABEL org.opencontainers.image.licenses='GPL v3'
LABEL org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.url="https://github.com/timo-reymann/SchemaNest"
LABEL org.opencontainers.image.documentation="https://github.com/timo-reymann/SchemaNest"
LABEL org.opencontainers.image.source="https://github.com/timo-reymann/SchemaNest.git"
RUN adduser -D -u 1000 schemanest
USER 1000
COPY --from=bin /bin/schemanest-cli /bin/schemanest-cli
