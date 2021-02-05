package utils

import (
	"context"
	"github.com/docker/docker/client"
	"net"
	"net/http"
	"os"
)

var Ctx context.Context
var Cli *client.Client
var Httpc http.Client
var systemConfig SystemConfig

type SystemConfig struct {
	Dir string
}

func init() {
	var err error

	Ctx = context.Background()
	Cli, err = client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	Httpc = http.Client{
		Transport: &http.Transport{
			DialContext: func(_ context.Context, _, _ string) (net.Conn, error) {
				return net.Dial("unix", "/var/run/docker.sock")
			},
		},
	}
	if err != nil {
		Logger().Error("初始化Docker上下文....FAIL!")
		return
	}

	homeDir, err := os.Hostname()
	if err != nil {
		Logger().Error("Home 目录获取失败....FAIL!")
		homeDir = "/tmp"
		return
	}
	homeDir = homeDir + "/.local/dobo"
	systemConfig.Dir = homeDir

	Logger().Info("初始化Docker上下文....OK!")
}

