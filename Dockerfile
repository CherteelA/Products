FROM mongo:7.0
RUN mkdir -p /etc/mongo
COPY mongod.conf /etc/mongo/

COPY init.js /docker-entrypoint-initdb.d/

RUN mkdir -p /backups

RUN chown mongodb:mongodb /docker-entrypoint-initdb.d/init.js

EXPOSE 27017
CMD ["mongod", "--config", "/etc/mongo/mongod.conf"]