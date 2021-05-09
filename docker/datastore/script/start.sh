#! /bin/sh

until mysqladmin ping -h mysql --silent; do
  echo 'waiting for mysql'
  sleep 2
done

echo "datastore is started!"
exec /main