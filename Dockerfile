FROM docker.io/library/golang:1.18 as builder

COPY / /noemptyhttpproxy
WORKDIR /noemptyhttpproxy
RUN CGO_ENABLED=0 make

FROM docker.io/library/golang:1.18
COPY --from=builder /noemptyhttpproxy/noemptyhttpproxy /usr/bin/
CMD ["noemptyhttpproxy"]
