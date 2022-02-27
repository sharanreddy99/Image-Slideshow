FROM php:7.4-apache

WORKDIR /var/www/html

RUN apt-get update && apt-get install zip netcat -y && \
    docker-php-ext-install pdo pdo_mysql && \
    docker-php-ext-install mysqli && \
    docker-php-ext-enable mysqli && \
    (curl -sS "https://getcomposer.org/installer" | php)

COPY compose* ./
RUN php composer.phar install

COPY ./start.sh ../
RUN ["chmod","+x","../start.sh"]

COPY . ./

EXPOSE 80

CMD ["../start.sh"]