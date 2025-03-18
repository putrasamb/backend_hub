FROM golang:1.22.9 AS build-stage

WORKDIR /app

COPY go.mod go.sum ./ 

RUN go mod download

COPY  . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /web cmd/web/main.go
# RUN CGO_ENABLED=0 GOOS=linux go build -o /worker cmd/worker/main.go

RUN mkdir -p /app/public/logs && \
    chmod 755 /app/public/logs 

# RUN  the tests in the container
# FROM build-stage AS run-test-stage
# RUN go test -v ./...


# Deploy the application binary into a lean image
FROM gcr.io/distroless/static-debian12 AS build-release-stage

WORKDIR /

COPY .enva /.env

# ARG ENV_FILE
# COPY ${ENV_FILE} /app/.env

COPY --from=build-stage /web /web
# COPY --from=build-stage /worker /worker
# COPY --from=build-stage /app/public /public

# RUN chmod 755 /public/logs
# #

EXPOSE ${PORT}

# USER nonroot:nonroot

ENTRYPOINT [ "/web" ]