FROM alpine:edge

# Set up dependencies
ENV PACKAGES go make git libc-dev bash

# Set up GOPATH & PATH

ENV GOPATH       /root/go
ENV BASE_PATH    $GOPATH/src/github.com/irisnet
ENV REPO_PATH    $BASE_PATH/iris-api-server
ENV WORKDIR      /irisnet/
ENV PATH         $GOPATH/bin:$PATH

# Link expected Go repo path

RUN mkdir -p $WORKDIR $GOPATH/pkg $ $GOPATH/bin $BASE_PATH

# Add source files

COPY . $REPO_PATH

# Install minimum necessary dependencies, build iris-api-server
RUN apk add --no-cache $PACKAGES && \
    cd $REPO_PATH && make all && \
    apk del $PACKAGES

# Set entrypoint

ENTRYPOINT ["iris-api"]