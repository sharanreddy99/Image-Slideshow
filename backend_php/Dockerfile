FROM php:7.4-apache

WORKDIR /var/www/html

RUN apt-get update && apt-get install zip -y

COPY . .

RUN docker-php-ext-install pdo pdo_mysql && \
    docker-php-ext-install mysqli && \
    docker-php-ext-enable mysqli 

RUN (curl -sS "https://getcomposer.org/installer" | php)


RUN php composer.phar install

EXPOSE 80

RUN ["chmod","+x","start.sh"]
CMD ["./start.sh"]