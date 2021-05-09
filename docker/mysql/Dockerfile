FROM mysql:5.7

RUN apt-get update -y && \
    apt-get install -y locales && \
    echo "ja_JP.UTF-8 UTF-8" > /etc/locale.gen && \
    locale-gen && \
    update-locale LANG=ja_JP.UTF-8

ENV LANG="ja_JP.UTF-8"

COPY ./my.cnf /etc/mysql/conf.d/my.cnf
COPY ./sql /docker-entrypoint-initdb.d
