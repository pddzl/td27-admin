FROM node:18

WORKDIR /web
COPY . .

RUN npm install pnpm --location=global --registry=http://registry.npm.taobao.org && pnpm config set registry http://registry.npm.taobao.org && pnpm i && pnpm build:prod

FROM nginx:alpine
LABEL MAINTAINER="pddzl5@foxmail.com"

RUN mkdir /usr/share/nginx/html/dist

COPY ./nginx/default.conf /etc/nginx/conf.d/default.template
COPY --from=0 /web/dist /usr/share/nginx/html/dist
