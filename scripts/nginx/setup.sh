#!/bin/bash

# shellcheck disable=SC2016
envsubst '${HTTPS_PORT} ${APP_HOST} ${APP_PORT}' < /etc/nginx/nginx.template.conf > /etc/nginx/conf.d/default.conf

nginx -g "daemon off;" -c /etc/nginx/nginx.conf