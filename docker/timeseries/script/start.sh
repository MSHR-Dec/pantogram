#! /bin/sh

while true; do
    code=$(curl -LI -XGET http://influxdb:8086/health -o /dev/null -w '%{http_code}' -s)
    if [ $code == 200 ]; then
      break
    fi
    echo 'waiting for influxdb'
    sleep 2
done

echo "timeseries is started!"
exec /main
