#**********************************************************************************************************************#
#
#     _____    ______   ______   __  __   ______   ______   ______  __   __       ______
#    /\  __-. /\  __ \ /\  ___\ /\ \/ /  /\  ___\ /\  == \ /\  ___\/\ \ /\ \     /\  ___\
#    \ \ \/\ \\ \ \/\ \\ \ \____\ \  _"-.\ \  __\ \ \  __< \ \  __\\ \ \\ \ \____\ \  __\
#     \ \____- \ \_____\\ \_____\\ \_\ \_\\ \_____\\ \_\ \_\\ \_\   \ \_\\ \_____\\ \_____\
#      \/____/  \/_____/ \/_____/ \/_/\/_/ \/_____/ \/_/ /_/ \/_/    \/_/ \/_____/ \/_____/
#
# Copyright (C) 2018 - 2021 J&J Ideenschmiede GmbH <info@jj-ideenschmiede.de>
#
#**********************************************************************************************************************#

# Load from docker
FROM golang:latest

# Config gin server
ENV ENABLE_SSL='true'
ENV CERTIFICATE_CRT_NAME='default'
ENV CERTIFICATE_KEY_NAME='default'

# Define port
EXPOSE 443

# Copy files to workdir
COPY . /go/src/app
WORKDIR /go/src/app

# Load libraries
RUN go get ./

# Build application
RUN go build -o app

# Start application
CMD ./app