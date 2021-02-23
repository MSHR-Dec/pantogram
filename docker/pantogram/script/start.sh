#! /bin/sh

until mysqladmin ping -h mysql --silent; do
  echo 'waiting for mysql'
  sleep 2
done

echo "pantogram is started!"
exec /main