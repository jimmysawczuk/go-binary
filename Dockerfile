FROM alpine
RUN apk add --no-cache tzdata
WORKDIR /home
COPY build/go-binary /usr/bin/go-binary
ENTRYPOINT ["go-binary"]
