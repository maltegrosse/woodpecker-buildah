FROM golang:1.21 AS builder
RUN mkdir -p /go/src/wrapper
WORKDIR /go/src/wrapper
COPY . ./
RUN go mod download
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=$ARCH go build -a -o /wrapper .

FROM docker.io/alpine:latest AS run

LABEL "io.containers.capabilities"="CHOWN,DAC_OVERRIDE,FOWNER,FSETID,KILL,NET_BIND_SERVICE,SETFCAP,SETGID,SETPCAP,SETUID,SYS_CHROOT"

RUN set -xeuf; \
  apk add --no-cache buildah fuse-overlayfs qemu; \
  adduser -D build; \
  echo -e "build:1:999\nbuild:1001:64535" > /etc/subuid; \
  echo -e "build:1:999\nbuild:1001:64535" > /etc/subgid; \
  :

COPY --from=builder --chmod=755 /wrapper /usr/local/bin/wrapper

COPY storage.conf containers.conf /etc/containers/

USER build

ENV BUILDAH_ISOLATION=chroot HOME=/home/build USER=build LOGNAME=build

ENTRYPOINT [ "/usr/local/bin/wrapper" ]