FROM nginx:1.17.4-alpine

COPY dist/  /usr/share/nginx/html/dist
RUN rm -rf /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80