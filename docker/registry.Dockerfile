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
    cp /dist/schema-nest-registry-linux-${architecture} /bin/schema-nest-registry && \
    chmod +x /bin/schema-nest-registry && \
    chown 1000:1000 /bin/schema-nest-registry

FROM chainguard/wolfi-base
LABEL org.opencontainers.image.title="schemanest-registry"
LABEL org.opencontainers.image.description="Registry for storing and managing schemas."
LABEL org.opencontainers.image.ref.name="main"
LABEL org.opencontainers.image.licenses='GPL v3'
LABEL org.opencontainers.image.vendor="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.authors="Timo Reymann <mail@timo-reymann.de>"
LABEL org.opencontainers.image.url="https://github.com/timo-reymann/SchemaNest"
LABEL org.opencontainers.image.documentation="https://github.com/timo-reymann/SchemaNest"
LABEL org.opencontainers.image.source="https://github.com/timo-reymann/SchemaNest.git"
COPY --from=license / /
RUN adduser -D -u 1000 schema-nest
USER 1000
COPY --from=bin /bin/schema-nest-registry /bin/schema-nest-registry
