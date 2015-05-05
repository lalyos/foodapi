# GOOS=linux go build && docker build -t food .
FROM progrium/busybox

RUN opkg-install curl bash

ENV VERSION 0.0.7
RUN curl -kLo /foodapi https://drone.io/github.com/lalyos/foodapi/files/foodapi \
   && chmod +x /foodapi
#COPY ./foodapi /

COPY start.sh /

EXPOSE 8080
CMD /start.sh
