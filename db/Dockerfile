FROM mysql

# ENV MYSQL_ROOT_PASSWORD <password>
ENV MYSQL_DATABASE fizzyFactory

COPY ./init.sql /docker-entrypoint-initdb.d/

EXPOSE 3306