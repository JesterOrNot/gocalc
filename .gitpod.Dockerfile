FROM gitpod/workspace-full-vnc
USER root
RUN apt-get update && apt-get install gnuplot-x11 r-base libgtk-3-dev libasound2-dev libnss3-dev
RUN go get github.com/Arafatk/glot
# Install custom tools, runtime, etc. using apt-get
# For example, the command below would install "bastet" - a command line tetris clone:
#
# RUN apt-get update \
#    && apt-get install -y bastet \
#    && apt-get clean && rm -rf /var/cache/apt/* && rm -rf /var/lib/apt/lists/* && rm -rf /tmp/*
#
# More information: https://www.gitpod.io/docs/42_config_docker/
