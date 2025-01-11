#!/bin/bash

BASEDIR="$(dirname $0)"
source "$BASEDIR/functions"

RANDOM_KEY="$(random_key)"
MY_IP="$(my_ip 6)"

setup_server
send_push
