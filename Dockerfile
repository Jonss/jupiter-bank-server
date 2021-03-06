FROM alpine:3.15.1 as alpine

COPY pkg/db/migrations/ pkg/db/migrations
COPY jupiterbank/bin app/bin
COPY app.env .

CMD ["app/bin"]