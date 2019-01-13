package main

import (
	"time"
)

const (
	Unknown      = "?"
	UpdatePeriod = time.Second
)

var Items = []statusFunc{
	netStatus("wlp4s0", "enp0s31f6"),
	batteryStatus("BAT0"),
	audioStatus(),
	memStatus,
	timeStatus,
}
