# base go image
FROM golang:1.21.3 as builder

ARG WORK_DIR=${WORK_DIR}

RUN mkdir -p ${WORK_DIR}

COPY . ${WORK_DIR}

WORKDIR ${WORK_DIR}

ARG GATEWAY_BINARY_FILE=${GATEWAY_BINARY_FILE}

RUN make server_build

RUN chmod +x ${WORK_DIR}/${GATEWAY_BINARY_FILE}

# build a tiny docker image
FROM alpine:3.18.4

ARG WORK_DIR=${WORK_DIR}

RUN mkdir -p ${WORK_DIR}

ARG GATEWAY_BINARY_FILE=${GATEWAY_BINARY_FILE}

COPY --from=builder ${WORK_DIR}/${GATEWAY_BINARY_FILE} ${WORK_DIR}
COPY --from=builder ${WORK_DIR}/app.env ${WORK_DIR}

ENV CONFIG_PATH=${CONFIG_PATH}

CMD [ "/var/www/app/gatewayApp" ]