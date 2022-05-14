module osiris/run

go 1.18

replace osiris/regmanip => ../regmanip

require osiris/regmanip v0.0.0-00010101000000-000000000000

require (
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/tklauser/go-sysconf v0.3.10 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
)

require (
	golang.org/x/sys v0.0.0-20220503163025-988cb79eb6c6 // indirect
	osiris/procexplorer v0.0.0-00010101000000-000000000000
)

replace osiris/procexplorer => ../procexplorer
