ARG VARIANT
FROM mcr.microsoft.com/vscode/devcontainers/go:${VARIANT}

ARG DEBIAN_FRONTEND noninteractive

RUN apt-get update \
 && apt-get install --no-install-recommends -y \
    libpcap-dev

ARG DEV_USER=vscode
