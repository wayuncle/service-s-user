package main

import (
	"flag"
	"fmt"
	"service-s-user/config/config"
	"service-s-user/handler"

	"github.com/micro/go-micro/v2/registry"
	"github.com/micro/go-micro/v2/web"
	"msp-git.connext.com.cn/connext-go-core/core-config/pconfig"
	"msp-git.connext.com.cn/connext-go-core/core-discovery/pregistry"
	"msp-git.connext.com.cn/connext-go-core/core-util/prouter"
	"msp-git.connext.com.cn/connext-go-core/core-util/prpc"
	"msp-git.connext.com.cn/connext-go-core/core-version/pversion"
	"msp-git.connext.com.cn/connext-go-third/third-log/plog"
)

func main() {
	// Init command flag
	ver := flag.Bool("version", false, "show version info")
	discoveryAddress := flag.String("regi_addr", "", "http://127.0.0.1:5214")
	flag.Parse()
	if *ver {
		pversion.FullVersion()
		return
	}
	// Init build version
	pversion.FullVersion()
	// Init Registry
	reg := pregistry.NewRegistry(registry.Addrs(func() string {
		if *discoveryAddress == "" {
			return pconfig.DiscoveryAddress()
		}
		return *discoveryAddress
	}()))
	// Init http client
	prpc.SetRegistry(reg)
	// Init confclient
	config.SetConfig(config.FIRSTTIME, "service-s-user")
	// Init DB config
	config.SetDB(config.FIRSTTIME)
	// Init service
	service := web.NewService(
		web.Name("service-s-user"),
		web.Address("0.0.0.0:"),
		web.Registry(reg),
		// web.RegisterTTL(pconsul.CRegisterTtL()),
		// web.RegisterInterval(pconsul.CRegisterInterval()),
	)
	if err := service.Init(); err != nil {
		plog.Error("service-s-user main", "%s", err.Error())
		return
	}

	// Init service's router
	handler.Wrapper()
	handler.Register(service)
	fmt.Println(prouter.PrintReport())

	// Service run
	if err := service.Run(); err != nil {
		plog.Error("service-s-user main", "%s", err.Error())
	}
}
