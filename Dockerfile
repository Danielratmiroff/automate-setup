# syntax=docker/dockerfile:1
# FROM golang:1.18-alpine
FROM ubuntu:latest

ENV DEBIAN_FRONTEND=noninteractive

# ARG DEBCONF_NOWARNINGS="yes" -- ignore pip sudo rights warning
USER root

# Install dependencies
RUN apt-get update && apt-get install -y \
    software-properties-common

RUN add-apt-repository universe

RUN apt-get update && apt-get install -y \
  python-is-python3 \
  python3-pip \
  openssh-client \
  # TODO: remove this -- only for dev purposes
  vim 

RUN pip install --upgrade pip \
  pip install ansible

COPY . ./ansible

WORKDIR /ansible 

EXPOSE 8080
