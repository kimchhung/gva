FROM node:20-alpine AS builder

WORKDIR /app
COPY /app/dist.tar.gz .
RUN tar -xzvf dist.tar.gz -C /app

FROM nginx:stable-alpine as production


RUN echo "types { application/javascript js mjs; }" > /etc/nginx/conf.d/mjs.conf
COPY --from=builder /app/dist /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/nginx.conf

EXPOSE 3000
