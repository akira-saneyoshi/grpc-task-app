FROM golang:1.23.6-alpine3.21
ARG PROTOBUF_VERSION=29.3
EXPOSE 8080

RUN apk update \
&& apk add alpine-sdk \
&& apk --update add tzdata \
&& cp /usr/share/zoneinfo/Asia/Tokyo /etc/localtime \
&& apk del tzdata \
&& rm -rf /var/cache/apk/*

WORKDIR /tmp/protoc
RUN apk add --no-cache unzip curl \
  && curl -L https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOBUF_VERSION}/protoc-${PROTOBUF_VERSION}-linux-x86_64.zip -o protoc.zip \
  && unzip protoc.zip \
  && mv /tmp/protoc/bin/* /usr/local/bin/ \
  && mv /tmp/protoc/include /usr/local/include \
  && rm -rf protoc protoc.zip \
  && apk add --no-cache protobuf-dev

WORKDIR /opt/grpc_task_backend
COPY ./backend /opt/grpc_task_backend
RUN go mod tidy