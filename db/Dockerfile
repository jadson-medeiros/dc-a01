FROM postgres:alpine

RUN mkdir -p /docker-entrypoint-initdb.d

COPY init-db.sql /docker-entrypoint-initdb.d/init-db.sql