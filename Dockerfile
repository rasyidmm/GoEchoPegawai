FROM ubuntu:18.04

MAINTAINER rasyidmm

RUN apt-get install -y python-software-properties

RUN add-apt-repository ppa:duh/golang
RUN apt-get update
RUN apt-get install -y golang