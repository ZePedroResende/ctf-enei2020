FROM library/postgres

ENV POSTGRES_USER docker
ENV POSTGRES_PASSWORD docker
COPY db/init.sql /docker-entrypoint-initdb.d/
