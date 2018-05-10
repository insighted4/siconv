FROM golang:1.10.1-stretch as builder
MAINTAINER Daniel Negri

RUN set -x \
    && apt-get update \
    && apt-get install -y build-essential ca-certificates git-core \
    && rm -rf /var/lib/apt/lists/*

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

RUN set -x \
    && go get github.com/AlekSi/gocov-xml \
    && go get github.com/axw/gocov/gocov \
    && go get github.com/golang/lint \
    && go get github.com/t-yuki/gocover-cobertura \
    && go get github.com/tebeka/go2xunit \
    && go get github.com/golang-migrate/migrate/cli \
    && go get github.com/lib/pq

COPY . /go/src/github.com/insighted4/siconv
WORKDIR /go/src/github.com/insighted4/siconv

RUN set -x \
    && make testall \
    && make release-binary \
    && mkdir -p /usr/share/siconv \
    && cp -r ./docs /usr/share/siconv/. \
    && cp -r ./release/bin /usr/share/siconv/. \
    && cp -r ./results /usr/share/siconv/. \
    && go build -tags 'postgres' -o /usr/local/bin/migrate github.com/golang-migrate/migrate/cli \
    && echo "Build complete."

# Release
FROM debian:jessie
MAINTAINER Daniel Negri

ENV GIN_MODE=release

RUN set -x \
    && apt-get update \
    && apt-get install -y ca-certificates \
    && rm -rf /var/lib/apt/lists/*

COPY --from=builder /usr/share/siconv /usr/share/siconv
RUN ln -s /usr/share/siconv/bin/server /usr/local/bin/siconv

COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

WORKDIR /usr/share/siconv
EXPOSE 8080
CMD ["siconv", "serve"]
