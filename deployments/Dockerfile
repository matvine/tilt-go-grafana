FROM golang:1.16
WORKDIR /app
ADD . .
RUN go install ./
ENTRYPOINT tilt-go-grafana