FROM golang:alpine3.18 as builder
LABEL authors="tutunak"

COPY . /app
WORKDIR /app
RUN go build -o feedcheck .

FROM alpine:3.18 as production
LABEL authors="tutunak"
COPY --from=builder /app/feedcheck /app/feedcheck
RUN addgroup -S feedcheck && adduser -S feedcheck -G feedcheck && \
    chown -R feedcheck:feedcheck /app
USER feedcheck
WORKDIR /app
CMD ["./feedcheck"]
