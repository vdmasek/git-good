FROM golang:latest as builder
WORKDIR /app
ADD . /app
RUN CGO_ENABLED=0 go build -o docker

FROM scratch
WORKDIR /app
COPY --from=builder /app/docker .
COPY --from=builder /app/static/* ./static/
COPY --from=builder /app/templates/* ./templates/
COPY --from=builder /app/app.env .

CMD ["./docker"]
