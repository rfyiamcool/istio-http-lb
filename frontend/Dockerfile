FROM alpine:3.9

MAINTAINER rfyiamcool

# qinghua mirror fast
RUN echo "https://mirror.tuna.tsinghua.edu.cn/alpine/v3.9/main/" > /etc/apk/repositories

RUN apk update \
        && apk upgrade \
        && apk add --no-cache bash \
        bash-doc \
        bash-completion \
        && rm -rf /var/cache/apk/* \
        && /bin/bash

EXPOSE 3001
ADD frontend /
CMD ["/frontend"]
