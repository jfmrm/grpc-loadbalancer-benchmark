FROM --platform=linux/amd64 debian:bookworm-slim

RUN apt-get update && \
    apt-get install -y wget

RUN wget https://github.com/bojand/ghz/releases/download/v0.108.0/ghz-linux-x86_64.tar.gz && \
    tar -xzf ghz-linux-x86_64.tar.gz && \
    chmod +x ghz-web && \
    rm ghz-linux-x86_64.tar.gz

EXPOSE 80

CMD ["./ghz-web"]
