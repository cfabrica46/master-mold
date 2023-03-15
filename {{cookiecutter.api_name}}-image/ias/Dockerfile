FROM hub.fif.tech/base/golang:1.16-latest AS build

WORKDIR /app
COPY . .
RUN make build

FROM hub.fif.tech/base/alpine:3.16
WORKDIR /app
COPY --from=build /app/build/bin/{{cookiecutter.api_name}} /app

USER 1000

ENTRYPOINT [ "/app/{{cookiecutter.api_name}}" ][]