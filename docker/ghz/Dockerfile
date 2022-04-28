FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y wget

RUN wget https://github.com/bojand/ghz/releases/download/v0.108.0/ghz-linux-x86_64.tar.gz -O /tmp/ghz.tar.gz

RUN tar -xzf /tmp/ghz.tar.gz -C /usr/local/bin

COPY helloworld .
COPY docker/ghz/*.json /ghz/config/

CMD ["ghz"]
