FROM ubuntu:18.04
MAINTAINER Ramon Tayag <ramon.tayag@gmail.com>

RUN apt-get update && \
  apt-get install -y ca-certificates curl && \
  # Keep size small
  apt-get clean && \
  rm -rf /var/lib/apt/lists/*
ENV APP_DIR=/app
ENV PATH=$APP_DIR:$PATH
RUN mkdir $APP_DIR
WORKDIR $APP_DIR

ENV BRIDGE_VERSION=v0.0.30
ENV BRIDGE_BASENAME=bridge-$BRIDGE_VERSION-linux-amd64
ENV BRIDGE_FILENAME=$BRIDGE_BASENAME.tar.gz

RUN curl -L https://github.com/stellar/bridge-server/releases/download/$BRIDGE_VERSION/$BRIDGE_FILENAME -o $BRIDGE_FILENAME && \
  tar zxf $BRIDGE_FILENAME && \
  mv $BRIDGE_BASENAME/bridge /usr/local/bin && \
  rm $BRIDGE_FILENAME && \
  rm -r $BRIDGE_BASENAME

ADD app $APP_DIR

CMD ["/app/entrypoint.sh", "bridge"]
