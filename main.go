package main

import (
	"log"

	"github.com/micro/go-micro"
	"github.com/mistrq/sort-srv/handler"
	proto "github.com/mistrq/sort-srv/proto"
)

func main() {
	service := micro.NewService(
		micro.Name("go.micro.srv.sort"),
		micro.Version("latest"),
	)

	service.Init()

	proto.RegisterSortHandler(service.Server(), new(handler.Sort))

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
