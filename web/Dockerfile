FROM node:17

WORKDIR /web
COPY . .

RUN npm install pnpm --location=global

RUN pnpm i && pnpm build:prod

FROM nginx:alpine
LABEL MAINTAINER="pddzl5@foxmail.com"

RUN mkdir /usr/share/nginx/html/dist

COPY ./nginx/default.conf /etc/nginx/conf.d/default.template
COPY --from=0 /web/dist /usr/share/nginx/html/dist
