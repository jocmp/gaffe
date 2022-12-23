package gaffe

import "time"

type Benchmark struct {
	ServerID   int64   `csv:"Server ID"`
	Sponsor    string  `csv:"Sponsor"`
	ServerName string  `csv:"Server Name"`
	Timestamp  string  `csv:"Timestamp"`
	Distance   string  `csv:"Distance"`
	Ping       float64 `csv:"Ping"`
	Download   float64 `csv:"Download"` // Download speed in bits
	Upload     float64 `csv:"Upload"`   // Upload speed in bits
	IPAddress  string  `csv:"IP Address"`
}

func (b *Benchmark) ChicagoTime() time.Time {
	location, _ := time.LoadLocation("America/Chicago")
	value, _ := time.Parse(time.RFC3339, b.Timestamp)

	return value.In(location)
}
