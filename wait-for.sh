#!/bin/sh

host="$1"
port="$2"
shift 2
cmd="$@"

until mysqladmin -h"$host" -P"$port" ping >/dev/null 2>&1; do
  
  sleep 1
done

>&2 echo "MySQL is up - executing command"
exec $cmd