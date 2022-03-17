FROM alpine:v3.15 as alpine

COPY jupiterbank/bin app/bin

CMD ["app/bin"]
