# Dockerfile
FROM postgres:12.20-alpine3.20

COPY ./sql-scripts/*.sql /docker-entrypoint-initdb.d/


