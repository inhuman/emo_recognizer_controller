FROM golang:1.18-alpine AS builder

ENV CGO_ENABLED 0

ENV TZ=Europe/Moscow

RUN apk --no-cache add ca-certificates tzdata && \
    cp -r -f /usr/share/zoneinfo/$TZ /etc/localtime

WORKDIR /app

COPY . .

RUN go build -mod=vendor -o /emo_recognizer_controller ./cmd/emo_recognizer_controller

COPY ./migrations /migrations
#
#FROM scratch
#
#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
#
#COPY --from=builder /etc/localtime /etc/localtime
#
#COPY --from=builder /emo_recognizer_controller /emo_recognizer_controller
#
#COPY --from=builder /migrations /migrations

ENTRYPOINT ["/emo_recognizer_controller"]

# http
EXPOSE 80