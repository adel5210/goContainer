package main

import (
	"fmt"
	"github.com/adel5210/goContainer/service/docker"
	"github.com/adel5210/goContainer/util/filereader"
)

func main() {
	dockerYML := filereader.ReadDockerYML("docker-container.yml")

	for k := range dockerYML {
		fmt.Println(dockerYML[k].Name)
		fmt.Println(dockerYML[k].Port)
		fmt.Println(dockerYML[k].Env)

		docker.CreateContainer(dockerYML[k].Name, dockerYML[k].Port)
	}

}
