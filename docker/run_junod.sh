#!/bin/sh

if test -n "$1"; then
    # need -R not -r to copy hidden files
    cp -R "$1/.bluechip" /root
fi

mkdir -p /root/log
bluechipd start --rpc.laddr tcp://0.0.0.0:26657 --trace