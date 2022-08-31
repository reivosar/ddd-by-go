FROM mysql:8.0
ENV LANG ja_JP.UTF-8

ADD ./conf/my.cnf /etc/mysql/conf.d/my.cnf

RUN mkdir /var/log/mysql
RUN touch /var/log/mysql/mysqld.log

CMD ["mysqld"]