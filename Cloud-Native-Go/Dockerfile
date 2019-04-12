# The base image is golang:1.12.3
FROM golang:1.12.3-alpine3.9
MAINTAINER Jieqiong YU

ENV SOURCES $HOME/go/src/github.com/go-pool/Cloud-Native-Go/

# Copy the source code directory to the base image
COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

# ENDPOINT is the thing which gets run when the docker container is running
ENTRYPOINT Cloud-Native-Go