FROM busybox:latest
RUN mkdir -p /app
COPY ./bin/mock-http-server /app
COPY ./config.json /app 
WORKDIR /app
CMD ["./mock-http-server","-config","/app/config.json","-addr",":8080"]
