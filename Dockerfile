FROM golang:1.18-alpine

RUN apk add --update --no-cache \
    make \
    git

WORKDIR /go/src/github.com/gandarez/check-pr-body-action

COPY . .

# build
RUN make build-linux

# apply permissions
RUN chmod a+x ./build/linux/amd64/check-pr-body-action

# symbolic link
RUN ln -s /go/src/github.com/gandarez/check-pr-body-action/build/linux/amd64/check-pr-body-action /bin/

# Specify the container's entrypoint as the action
ENTRYPOINT ["/bin/check-pr-body-action"]
