package SkuPi

type Pi struct {
	ConnId     int64
	Info       *Info
	IsTimeSync bool
}

type Info struct {
	Name       string
	ConnectNow int
}
