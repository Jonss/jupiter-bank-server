FROM alpine as alpine

COPY jupiterbank/bin app/bin

CMD gunicorn --bind 0.0.0.0:$PORT wsgi
ENTRYPOINT ["app/bin -p $PORT"] 
