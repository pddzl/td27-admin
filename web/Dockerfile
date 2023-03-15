FROM node:17

WORKDIR /web
COPY . .

RUN pnpm -i && pnpm build

FROM nginx:alpine
LABEL MAINTAINER="pddzl5@foxmail.com"

COPY ./nginx/default.conf /etc/nginx/conf.d/default.template
COPY --from=0 /web/dist /usr/share/nginx/html
