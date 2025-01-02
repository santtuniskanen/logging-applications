package main

import (
	"fmt"
	"net"
	"time"

	"github.com/BurntSushi/toml"
	"go.uber.org/zap"
)

type Config struct {
	Server struct {
		Host         string        `toml:"host"`
		Port         int           `toml:"port"`
		ReadTimeout  time.Duration `toml:"read_timeout"`
		WriteTimeout time.Duration `toml:"write_timeout"`
	} `toml:"server"`
}

func loadConfig(path string) (*Config, error) {
	var config Config
	_, err := toml.DecodeFile(path, &config)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}
	return &config, nil
}

var logger *zap.Logger

func initLogger() {
	logger := zap.Must(zap.NewProduction())
}

func listener() error {
	listener, err := net.Listen("tcp", "127.0.0.1")
	if err != nil {
		logger.Info("Could not listen to address")
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func(c net.Conn) {
			defer c.Close()

		}(conn)
	}
	return nil
}

func main() {
	logger := zap.Must(zap.NewProduction())
	defer logger.Sync()
	logger.Info("Hello, world!")
}
