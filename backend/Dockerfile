FROM alpine

RUN apk update && \ 
    apk add \ 
      ca-certificates && \
    rm -rf /var/cache/apk/*

EXPOSE 3000
ADD backend /
CMD ["/backend"]
