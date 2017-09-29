FROM golang:1.6-alpine

ENV APP_DIR $GOPATH/src/github.com/yshuman1/gourl

ADD . $APP_DIR
WORKDIR $APP_DIR

RUN apk --update add git && rm -rf /var/cache/apk/*

# Install packages.
# Ideally we'd use `godep` for this, but to keep this short
# we'll just install them manually here.
RUN go get -u github.com/kataras/iris \
    go get -u gopkg.in/redis.v4 \
    go get -u github.com/codegangsta/gin

EXPOSE 3000
CMD gin run