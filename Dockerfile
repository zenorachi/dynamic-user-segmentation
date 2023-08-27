FROM alpine:latest

WORKDIR /root

# Copy binary
COPY ./.bin/app /root/.bin/app

# Copy configs
COPY ./.env /root/
COPY ./configs/main.yml /root/configs/main.yml

# Copy docs
COPY ./docs /root/docs

# Install psql-client
RUN apk add --no-cache postgresql-client

CMD ["sh", "-c", "sh ./scripts/db-connection/wait-db.sh && ./.bin/app"]