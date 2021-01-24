#! /bin/bash
set -Ceuo pipefail

SCRIPT_DIR=$(cd $(dirname $0); pwd)
DATA_DIR="$SCRIPT_DIR"/data
META_DATA="$DATA_DIR"/base.tsv

export MYSQL_PWD=pantogram

# curl -sSL https://rti-giken.jp/fhc/api/train_tetsudo/train.tsv -o "$META_DATA"

# company
coms=$(< "$META_DATA" cut -f1 | sort | uniq)
IFS=$'\n'
if [ ! -e "$DATA_DIR"/companies.text ]; then
  for c in $(echo "$coms"); do
    echo "('$c')," >> "$DATA_DIR"/companies.text
  done
fi

# line
if [ ! -e "$DATA_DIR"/routes.text ]; then
  while read -r line; do
    com=$(echo "$line" | cut -f1)
    l=$(echo "$line" | cut -f2)
    id=$(mysql -upantogram -h127.0.0.1 -P53306 pantogram -e "select id from companies where name = '$com';" | cut -d$'\n' -f2)
    echo "('$l', $id)," >> "$DATA_DIR"/routes.text
  done < "$META_DATA"
fi

# route_prefecture
r_id=$(mysql -upantogram -h127.0.0.1 -P53306 pantogram -e "select id from routes where name = '$1';" | cut -d$'\n' -f2)
p_id=$(mysql -upantogram -h127.0.0.1 -P53306 pantogram -e "select id from prefectures where name = '$2';" | cut -d$'\n' -f2)
echo "($r_id, $p_id)," >> "$DATA_DIR"/routes_prefectures.text
