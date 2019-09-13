# docker build -t noah-explorer-api -f DOCKER/Dockerfile .
# docker run -d -p
FROM golang:1.12-buster as builder

ENV APP_PATH /home/noah-explorer

COPY . ${APP_PATH}

WORKDIR ${APP_PATH}

RUN make create_vendor && \
    make build && \
    cp ./builds/explorer /usr/local/bin/explorer

CMD ["explorer"]
EXPOSE 9050
STOPSIGNAL SIGTERM