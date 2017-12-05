FROM alpine

COPY er-task /app/
COPY _template /app/_template/
COPY _static/js /app/_static/js/


RUN adduser -D -u 1000 user 
    

USER user
WORKDIR /app

ENTRYPOINT ["/app/er-task"]