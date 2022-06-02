FROM golang:1.18.3-bullseye

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY *.go ./
RUN go build -o /indiepkg

COPY scripts/docker/portal.sh /
RUN chmod 755 /portal.sh

RUN useradd -m user
USER user

ENTRYPOINT ["/portal.sh"]
