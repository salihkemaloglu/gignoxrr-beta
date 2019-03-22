FROM golang:1.8

# expose default port
# EXPOSE 8000
# install protobuf from source
RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool
RUN git clone https://github.com/google/protobuf.git && \
    cd protobuf && \
    ./autogen.sh && \
    ./configure && \
    make && \
    make install && \
    ldconfig && \
    make clean && \
    cd .. && \
    rm -r protobuf
# Get the source from GitHub
RUN go get google.golang.org/grpc
# Install protoc-gen-go
RUN go get github.com/golang/protobuf/protoc-gen-go    
RUN go get gopkg.in/mgo.v2/bson
RUN go get github.com/rs/cors
RUN go get github.com/go-chi/chi

# RUN go get google.golang.org/grpc/codes
# RUN go get google.golang.org/grpc/status
# RUN go get google.golang.org/grpc/grpclog
# RUN go get github.com/go-chi/chi/middleware
RUN go get github.com/improbable-eng/grpc-web/go/grpcweb
RUN go get github.com/dgrijalva/jwt-go

# set environment path
ENV PATH /go/bin:$PATH

# cd into the api code directory
WORKDIR /go/src/github.com/salihkemaloglu/gignox-rr-beta-001

# create ssh directory
RUN mkdir ~/.ssh
RUN touch ~/.ssh/known_hosts
RUN ssh-keyscan -t rsa github.com >> ~/.ssh/known_hosts

# allow private repo pull
RUN git config --global url."https://e4d5159cc774d99744024453431f00ddbb8d7b1d:x-oauth-basic@github.com/".insteadOf "https://github.com/"

# copy the local package files to the container's workspace
ADD . /go/src/github.com/salihkemaloglu/go-docker

# install the program
RUN go install github.com/salihkemaloglu/go-docker

# start application
CMD ["go","run","main.go"] 