FROM debian:buster-slim

ENV APP_ROOT=/usr/app

ARG USER=golang
ARG GROUP=golang

WORKDIR $APP_ROOT

RUN apt update && apt install build-essential ca-certificates -y

WORKDIR $APP_ROOT

RUN groupadd -g 1001 $GROUP && \
    adduser --system --no-create-home --uid 1001 --ingroup $GROUP --disabled-password $USER && \
    chown 1001:1001 $APP_ROOT

USER $USER

COPY --chown=1001:1001 main .
#COPY --chown=1001:1001 application.yml .

EXPOSE 8080

CMD ["./main"]