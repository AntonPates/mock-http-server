FROM busybox:latest
RUN mkdir -p /app
COPY ./bin/mock-server /app
COPY ./config.json /app 
WORKDIR /app
CMD ["./mock-server","-config","/app/config.json","-addr",":8080"]
