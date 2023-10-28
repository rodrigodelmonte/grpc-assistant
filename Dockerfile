FROM golang:1.20-alpine@sha256:b9ff7091cb0888abc44844d581be7a15f6913c54253d84f4385164353a30d6a7 as builder

WORKDIR /
RUN apk --update add ca-certificates
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /bin/app

FROM scratch
COPY --from=builder /bin/app /bin/app
COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs/
ENTRYPOINT ["/bin/app"]
