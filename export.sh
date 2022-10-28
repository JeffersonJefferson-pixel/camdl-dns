#!/bin/sh

./dns.exe

devp2p dns sign ${1:+--domain "$1"} ${2:+--seq "$2"} data/ data/dnskey.json

devp2p dns to-txt data/ data/TXT.json