#!/bin/bash

# requires go, postgresql, bash, and curl to be installed locally


# start up the postgres

# for windows
pg_ctl -D "C:\Program Files\pgsql\data" start

# for mac, if postgres was downloaded via homebrew
# brew services start postgresql


# start up gokey-cURL, could also use "go run server.go"
air
