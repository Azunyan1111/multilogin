FROM mysql:5.7

LABEL architecture='mysql_ml'

# setting
COPY my.cnf /etc/mysql/conf.d/

RUN chmod 600 /etc/mysql/conf.d/my.cnf

# docker-entrypoint-initdb.dの中にある.sqlを勝手に実行してくれる
COPY init/intiSql.sql /docker-entrypoint-initdb.d/initSql.sql