# build stage
FROM golang:1.13.2-alpine AS go-build-env

RUN apk add --no-cache --update make git alpine-sdk gcc build-base

WORKDIR /src
ADD . /src

RUN make test
RUN make build_static

# run stage
FROM alpine 
# This App needs timezone info to work properly!
RUN apk --no-cache add tzdata
WORKDIR /app
COPY --from=go-build-env /src/bin/solidfire-exporter /app/
ENTRYPOINT ["./solidfire-exporter"]