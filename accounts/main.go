package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/laidingqing/KeepJobs/accounts/service"
	"github.com/laidingqing/KeepJobs/common/config"
	grpclb "github.com/laidingqing/KeepJobs/common/registry"
	pb "github.com/laidingqing/KeepJobs/pb"
	"google.golang.org/grpc"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

func main() {
	config.LoadDefaults()
	config.ParseCmdParams(config.DefaultCmdLine{
		HostName:         "localhost",
		Port:             4100,
		ServiceName:      config.ServiceAccountName,
		RegistryLocation: "http://127.0.0.1:2379",
	})

	// Set up the core logger
	log.SetOutput(&lumberjack.Logger{
		Filename:   "./logs/accounts_service.log",
		MaxSize:    config.Logger.MaxSize,
		MaxBackups: config.Logger.MaxBackups,
		MaxAge:     config.Logger.MaxAge,
	})

	lis, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.Service.Port))
	if err != nil {
		panic(err)
	}

	err = grpclb.Register(config.Service.ServiceName,
		config.Service.DomainName,
		config.Service.Port,
		config.Service.RegistryLocation,
		time.Second*10,
		15)
	if err != nil {
		panic(err)
	}

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL, syscall.SIGHUP, syscall.SIGQUIT)
	go func() {
		s := <-ch
		log.Printf("receive signal '%v'", s)
		grpclb.UnRegister()
		os.Exit(1)
	}()

	log.Printf("starting user service at %d", config.Service.Port)
	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &service.RPCAccountServer{})
	s.Serve(lis)
}
