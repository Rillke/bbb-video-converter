FROM golang:1.23-alpine3.20 AS builder
ENV USER=appuser
ENV UID=10001
RUN apk update && apk add --no-cache git ca-certificates tzdata && update-ca-certificates
RUN mkdir -p /srv/bbb-convert /srv/bbb-convert/release
WORKDIR /srv/bbb-convert
COPY . .
RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags='-w -s -extldflags "-static"' -a \
        -o /srv/bbb-convert/release/bbb-convert

FROM python:3.9-alpine3.20
RUN apk add --no-cache git ca-certificates chromium ffmpeg
RUN adduser -D bigbluebutton bigbluebutton
COPY --from=builder /srv/bbb-convert/release /srv/bbb-convert
#USER bigbluebutton
WORKDIR /home/bigbluebutton
RUN cd /opt && git clone https://github.com/fossasia/bbb-download.git && cd bbb-download && pip install -r requirements.txt
# ENTRYPOINT ["/srv/bbb-convert/bbb-convert"]

