# build stage
FROM golang:alpine AS build-env
ADD . /src
RUN cd /src && go build -o goapp

# final stage
FROM alpine
RUN apk update && apk add flite
WORKDIR /app
COPY --from=build-env /src/goapp /app/
ENTRYPOINT ./goapp