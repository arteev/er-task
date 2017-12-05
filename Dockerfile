FROM alpine

COPY er-task /app/

RUN adduser -D -u 1000 user

USER user
WORKDIR /app

ENTRYPOINT ["/app/er-task"]