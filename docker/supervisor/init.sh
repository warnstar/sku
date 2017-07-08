#!/usr/bin/env bash

chmod -R 777 /var/www/html/publish/develop/public
chmod -R 777 /var/www/html/publish/prod/public

chmod +x /var/www/html/publish/develop/app
chmod +x /var/www/html/publish/prod/app