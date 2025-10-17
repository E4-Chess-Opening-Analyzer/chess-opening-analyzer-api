FROM mongo:7-jammy
LABEL maintainer="Luc Lacotte <luc.lonqq@gmail.fr>"
LABEL description="mongo:7-jammy for Chess Opening Analyzer"
LABEL version="1.0"

# Copy initialization script
COPY build/database/init-mongo.js /docker-entrypoint-initdb.d/