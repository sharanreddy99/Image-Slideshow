#!/bin/sh
echo "Waiting for mysql. This may take few more seconds ..."
until nc -z -v -w30 db 3306
do
    >&2 echo "Waiting for mysql. This may take few more seconds ..."
    sleep 10
done
echo "mysql connection established..."

printf '\n\nPHP APIs can be accessed at the following url: '
tail -n 1 /etc/hosts | awk '{printf $1;}'
printf ':'${IMAGE_VIEWER_PHP_PORT}
printf '\n---------------------------------------------------------------------\n'

printf 'PHPMYADMIN can be accessed at the following url: localhost:'${IMAGE_VIEWER_PHPMYADMIN_PORT}
printf '\n---------------------------------------------------------------------\n'

# Start server
/usr/local/bin/apache2-foreground