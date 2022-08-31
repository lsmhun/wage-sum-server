FROM alpine
WORKDIR /app
COPY out/bin/wagesum /app

CMD ["/app/wagesum"]