FROM alpine

COPY er-task /app/
COPY src/templates /app/templates/
COPY src/static/js /app/static/js/


RUN adduser -D -u 1000 user 
    

USER user
WORKDIR /app

ENTRYPOINT ["/app/er-task"]
