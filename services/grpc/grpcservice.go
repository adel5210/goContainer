package grpc

import (
	"context"
	"github.com/adel5210/goContainer/services/someclient"
	"google.golang.org/grpc"
	"log"
	"os"
)

func init(){
	host:=os.Getenv("CLIENT_HOST")

	conn, err := grpc.Dial(host, grpc.WithInsecure())
	if nil != err{
		log.Fatal(err.Error())
	}

	defer conn.Close()

	client := someclient.NewClientServiceClient(conn)
	name:="adel"
	add:="home"
	request := &someclient.ClientRequest{
		Name: &name,
		Address: &add,
	}

	response ,err := client.SomeClientFunction(context.Background(), request)
	if nil != err{
		log.Fatal(err.Error())
	}

	log.Println(response.Message)

}
