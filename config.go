package server

import "time"

// Config values for the server
type Config struct {
	ListenAddress    string        `kong:"default=':8080',help='Address on which the service will run',short='l',env='LISTEN_ADDR'"`
	ShutdownDuration time.Duration `kong:"default='1s',help='Duration to take to allow existing connections to complete',short='s'"`
	WarningDuration  time.Duration `kong:"default='1s',help='Duration to take to allow routing algorithms to reroute',short='w'"`
}
