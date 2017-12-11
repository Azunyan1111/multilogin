FROM mysql:5.7

LABEL architecture='mysql_ml'

# setting
COPY my.cnf /etc/mysql

ENV MYSQL_ROOT_PASSWORD = password

RUN chmod 600 /etc/mysql/my.cnf

# docker-entrypoint-initdb.dの中にある.sqlを勝手に実行してくれる
COPY intiSql.sql /docker-entrypoint-initdb.d/initSql.sql
