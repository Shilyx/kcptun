module github.com/xtaci/kcptun

require (
	github.com/coreos/go-iptables v0.4.2 // indirect
	github.com/golang/snappy v0.0.1
	github.com/google/gopacket v1.1.17 // indirect
	github.com/pkg/errors v0.8.1
	github.com/templexxx/cpufeat v0.0.0-20180724012125-cef66df7f161 // indirect
	github.com/urfave/cli v1.21.0
	github.com/xtaci/kcp-go v5.4.19+incompatible
	github.com/xtaci/lossyconn v0.0.0-20190602105132-8df528c0c9ae // indirect
	github.com/xtaci/smux v1.4.6
	github.com/xtaci/smux/v2 v2.0.16
	github.com/xtaci/tcpraw v1.2.25
	golang.org/x/crypto v0.0.0-20191206172530-e9b2fee46413
	golang.org/x/sys v0.0.0-20190910064555-bbd175535a8b // indirect
)

replace github.com/xtaci/kcp-go => ./kcp-go

go 1.13
