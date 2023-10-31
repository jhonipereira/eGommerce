# build docker image
FROM alpine:latest

RUN mkdir /app

COPY prodApp /app

CMD [ "/app/prodApp" ]