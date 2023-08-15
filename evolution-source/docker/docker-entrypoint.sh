#!/bin/bash
set -e

echo "Container's IP address: `awk 'END{print $1}' /etc/hosts`"
echo "Working dir: `pwd`"

cd /opt/evolution-source/src/
go build -o service main.go
./service
exec "$@"