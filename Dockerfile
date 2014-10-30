FROM debian:jessie

RUN apt-get update && apt-get install -y curl
RUN curl -Lo /foodapi https://drone.io/github.com/lalyos/foodapi/files/foodapi \
  && chmod +x /foodapi

COPY start.sh /
CMD /start.sh
