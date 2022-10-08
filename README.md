# goIPASN
https://github.com/hadiasghari/pyasn

`cd ~/asn/;pyasn_util_asnames.py > asn.json;pyasn_util_download.py -4;rib=$(ls -t rib.*.bz2 | head -1);pyasn_util_convert.py --single $rib IPASN.DAT; rm -f $rib`
`echo 8.8.8.8 | asnlook`
`8.8.8.8 15169 8.8.8.0/24 GOOGLE, US`
