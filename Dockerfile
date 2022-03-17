FROM alpine as alpine

COPY jupiterbank/bin app/bin

EXPOSE 8080
CMD ["app/bin"]
