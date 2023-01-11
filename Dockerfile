FROM debian:buster-slim

ENV WAGESUM_DB_TYPE=postgres
ENV WAGESUM_DB_HOST=127.0.0.1
ENV WAGESUM_DB_PORT=5432
ENV WAGESUM_DB_NAME=wagesum
ENV WAGESUM_DB_USERNAME=postgres
ENV WAGESUM_HTTP_SERVER_PORT=3000
ENV WAGESUM_DB_PASSWORD=

WORKDIR /app
ADD out/bin/wagesum /app/wagesum
COPY out/bin/wagesum /app/wagesum

CMD ["/app/wagesum"]