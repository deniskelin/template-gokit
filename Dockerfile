# DOCKER_BUILDKIT=0 docker build -t test-protoc --progress plain .

##########################
#                        #
# BUILD PACKER CONTAINER #
#                        #
##########################

#FROM alpine as packer
#
#COPY builded_app /builded_app
#
#RUN apk add --no-cache upx && \
#    upx --lzma --best /builded_app && \
#    apk del upx


#########################
#                       #
# BUILD FINAL CONTAINER #
#                       #
#########################

FROM alpine

#COPY --from=packer /builded_app /usr/local/bin/ipcr-echo
COPY builded_app /usr/local/bin/rds

RUN set -exu && \
    env && \
    chmod -R a+x /usr/local/bin/* && \
    ls -lah /usr/local/bin

ENV HTTP_LISTEN_ADDRESS=":8080"

CMD rds

ARG BUILD_DATE
ARG LABEL_NAME
ARG LABEL_DESCRIPTION
ARG LABEL_VENDOR
ARG LABEL_URL
ARG LABEL_URL_SOURCE
ARG LABEL_URL_DOCUMENTATION
ARG LABEL_MAINTAINER

ARG CI_COMMIT_SHA
ARG CI_COMMIT_REF_NAME
ARG CI_COMMIT_MESSAGE
ARG CI_COMMIT_AUTHOR

MAINTAINER [riftbit] ErgoZ <https://riftbit.com/>

LABEL \
    org.opencontainers.image.version="${VERSION}" \
    org.opencontainers.image.created="${BUILD_DATE}" \
    org.opencontainers.image.revision="${CI_COMMIT_SHA}" \
    org.opencontainers.image.title="${LABEL_NAME}" \
    org.opencontainers.image.description="${LABEL_DESCRIPTION}" \
    org.opencontainers.image.vendor="${LABEL_VENDOR}" \
    org.opencontainers.image.authors="${LABEL_MAINTAINER}" \
    org.opencontainers.image.url="${LABEL_URL}" \
    org.opencontainers.image.source="${LABEL_URL_SOURCE}" \
    org.opencontainers.image.documentation="${LABEL_URL_DOCUMENTATION}" \
	Description="${LABEL_DESCRIPTION}" \
	maintainer="${LABEL_MAINTAINER}"