#!/bin/sh
printf '\n\nPHP APIs can be accessed at the following url: '
tail -n 1 /etc/hosts | awk '{printf $1;}'
printf ':'${IMAGE_VIEWER_PHP_PORT}
printf '\n---------------------------------------------------------------------\n'

printf 'PHPMYADMIN can be accessed at the following url: localhost:'${IMAGE_VIEWER_PHPMYADMIN_PORT}
printf '\n---------------------------------------------------------------------\n'

# Start server
/usr/local/bin/apache2-foreground