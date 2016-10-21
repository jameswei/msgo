package broker

import (
	"flag"
	"fmt"
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"math"
)

type config struct {
	HttpPort int
	MsgPort  int
	Auth     bool
	UserName string
	Token    string
	//Cluster name identifies your cluster for auto-discovery. If you're running
	//multiple clusters on the same network, make sure you're using unique names.
	//
	Retry int
}

var Config *config = new(config)

func LoadConfig() {
	var configFile string
	// todo the default config Path
	flag.IntVar(&Config.HttpPort, "httpport", 13000, "the http port")
	flag.IntVar(&Config.MsgPort, "port", 13001, "the msg port")
	flag.IntVar(&Config.Retry, "r", 3, "the retry times")
	flag.StringVar(&configFile, "c", "", "the config file path")
	flag.Parse()
	if configFile != "" {
		Bytes, err := ioutil.ReadFile(configFile)
		if err != nil {
			panic(fmt.Sprintf("reading config file error %s: %v", configFile, err))
		}
		if _, err := toml.Decode(string(Bytes), Config); err != nil {
			panic(fmt.Sprintf("parse config file error %s: %v", configFile, err))
		}
	}
	ArbitrateConfigs(Config)
}

func ArbitrateConfigs(c *config) {
	// check the ClusterName, ClusterName is used to Identify the clusters in the Local NetWork
	if Config.HttpPort == Config.MsgPort {
		panic("port conflict")
	}
	if Config.HttpPort > math.MaxInt16 || Config.HttpPort < 1024 {
		panic(fmt.Errorf("illegal http port %s", Config.HttpPort))
	}

	if Config.MsgPort > math.MaxInt16 || Config.MsgPort < 1024 {
		panic(fmt.Errorf("illegal msg port %s", Config.MsgPort))
	}

	if Config.Retry > 10 {
		Config.Retry = 10
	}
	if Config.Retry < 1 {
		Config.Retry = 1
	}
	Log.Printf("Message service listening on port :%d\n", Config.MsgPort)
}
