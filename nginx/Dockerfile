FROM nginx

COPY ./nginx.conf /etc/nginx/nginx.conf
RUN rm -r /etc/nginx/conf.d
COPY ./conf.d /etc/nginx/conf.d
COPY ./add_forwarded /etc/nginx/add_forwarded