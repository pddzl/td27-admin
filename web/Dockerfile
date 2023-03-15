FROM node:17

WORKDIR /web
COPY . .

RUN curl -f curl -f https://get.pnpm.io/v6.16.js | node - add --global pnpm

RUN pnpm i && pnpm build:prod

FROM nginx:alpine
LABEL MAINTAINER="pddzl5@foxmail.com"

RUN mkdir /usr/share/nginx/html/dist

COPY ./nginx/default.conf /etc/nginx/conf.d/default.template
COPY --from=0 /web/dist /usr/share/nginx/html/dist
