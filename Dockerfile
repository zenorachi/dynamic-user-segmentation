FROM alpine:latest

WORKDIR /root

# Copy binary
COPY ./.bin/app /root/.bin/app

# Copy migrations
COPY ./scripts/migrations /root/scripts/migrations

# Copy configs
COPY ./.env /root/
COPY ./configs/main.yml /root/configs/main.yml

# Copy wait-db.sh
COPY ./scripts/db-connection/wait-db.sh /root/scripts/db-connection/

# Copy secrets
COPY ./secrets/ /root/secrets/

# Copy docs
#COPY ./docs /root/docs

# Install psql-client
RUN apk add --no-cache postgresql-client

CMD ["sh", "-c", "sh ./scripts/db-connection/wait-db.sh && ./.bin/app"]