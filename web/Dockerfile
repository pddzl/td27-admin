# syntax=docker/dockerfile:1
FROM node:22-alpine AS builder

WORKDIR /web

# Enable pnpm via corepack (comes with Node)
RUN corepack enable && corepack prepare pnpm@9 --activate

# Copy dependency files first (better layer caching)
COPY package.json pnpm-lock.yaml ./
RUN pnpm install --frozen-lockfile

# Copy source and build
COPY . .
RUN pnpm build:prod

FROM nginx:1.27-alpine

LABEL maintainer="pddzl5@foxmail.com"

# Security: non-root user
RUN addgroup -g 1000 -S td27 && adduser -u 1000 -S td27 -G td27

COPY ./nginx/default.conf /etc/nginx/conf.d/default.template
COPY --from=builder /web/dist /usr/share/nginx/html/dist

RUN chown -R td27:td27 /usr/share/nginx/html/dist

# Healthcheck for nginx
HEALTHCHECK --interval=30s --timeout=5s --start-period=5s --retries=3 \
  CMD wget --quiet --tries=1 --spider http://localhost/ || exit 1

EXPOSE 80

CMD ["/bin/sh", "-c", "envsubst '$$BACKEND' < /etc/nginx/conf.d/default.template > /etc/nginx/conf.d/default.conf && exec nginx -g 'daemon off;'"]
