FROM alpine:3.17

ENV PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
EXPOSE 8080
COPY server /server
CMD ["/bin/sh", "-c", "/server"]

