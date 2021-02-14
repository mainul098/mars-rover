FROM golang:1.14

ENV APP_DIR /mars-rover
COPY . ${APP_DIR}
WORKDIR ${APP_DIR}

RUN make deps