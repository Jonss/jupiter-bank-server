FROM alpine:3.15.1 as alpine

COPY jupiterbank/bin app/bin

CMD ["app/bin"]
