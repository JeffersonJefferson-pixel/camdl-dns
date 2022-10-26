./dns.exe

devp2p dns sign --domain $1 --seq $2 data/ data/dnskey.json

devp2p dns to-txt data/ data/TXT.json