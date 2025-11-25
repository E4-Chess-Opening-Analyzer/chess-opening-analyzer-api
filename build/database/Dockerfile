FROM mongo:7-jammy
LABEL maintainer="Luc Lacotte <luc.lonqq@gmail.fr>"
LABEL description="mongo:7-jammy for Chess Opening Analyzer"
LABEL version="1.0"

COPY build/database/init-mongo.js.template /docker-entrypoint-initdb.d/init-mongo.js.template
COPY build/database/generate-init-mongo.sh /docker-entrypoint-initdb.d/generate-init-mongo.sh

RUN chmod +x /docker-entrypoint-initdb.d/generate-init-mongo.sh