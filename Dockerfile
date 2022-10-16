FROM --platform=linux/amd64 golang:1.21rc1-buster
WORKDIR /
ARG BINARY_NAME
ENV BINARY_NAME=$BINARY_NAME
COPY $BINARY_NAME .
RUN chmod 755 secure-cli-linux-amd64
ENTRYPOINT ["./secure-cli-linux-amd64"]