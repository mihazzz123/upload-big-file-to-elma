FROM golang:1.17.6-alpine3.14
WORKDIR /go/src/git.elewise.com/elma365/upload-big-file-elma365
RUN apk add --no-cache make git
COPY . .

ARG GOOS="linux"
ARG GOARCH="amd64"

RUN GOOS=${GOOS} GOARCH=${GOARCH} make build

# Фаза упаковки без лишних зависимостей
FROM dreg.elewise.com:5005/docker/base:latest
WORKDIR /srv
COPY data data
COPY --from=0 /go/src/git.elewise.com/elma365/upload-big-file-elma365/build/upload-big-file-elma365 /srv/bin/upload-big-file-elma365
RUN chown -R app:app /srv
EXPOSE 5000
EXPOSE 3000
CMD ["/srv/bin/upload-big-file-elma365"]