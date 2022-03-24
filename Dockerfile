ARG GO_VERSION=1.17.8
ARG ALPINE_VERSION=3.15.2

FROM --platform=$BUILDPLATFORM golang:${GO_VERSION}-alpine as builder
ARG TARGETOS
ARG TARGETARCH

# Download dependencies (this is done in a seperate layer to take advantage
# of Docker layer caching to decrease build times)
WORKDIR /src
COPY ./go.mod ./go.sum ./
RUN go mod download

# Build
COPY ./ /src
RUN GOOS=$TARGETOS GOARCH=$TARGETARCH CGO_ENABLED=0 go build -a -tags netgo -ldflags '-w' -o meetup-reminder

# Use Alpine for base image
FROM alpine:${ALPINE_VERSION}

# ENV SLACK_WEBHOOK https://hooks.slack.com/services/xxxxxxxxx/xxxxxxxxx/xxxxxxxxxxxxxxxxxxxx
# ENV MEETUP_NAME meetup_url_name
# ENV DAYS_BEFORE_REMINDER 7
# ENV MESSAGE "is in a week"docker-buildx

# Install binary
COPY --from=builder /src/meetup-reminder /usr/local/bin/meetup-reminder
ENTRYPOINT ["meetup-reminder"]
