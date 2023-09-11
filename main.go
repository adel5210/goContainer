package main

import (
	"fmt"
	"github.com/adel5210/goContainer/service/docker"
	"github.com/adel5210/goContainer/util/filereader"
)

func main() {
	dockerYML := filereader.ReadDockerYML("docker-container.yml")
	fmt.Println(dockerYML["MSSQL"].Name)
	fmt.Println(dockerYML["MSSQL"].Port)
	fmt.Println(dockerYML["MSSQL"].Env)

	docker.CreateContainer(dockerYML["MSSQL"].Name, dockerYML["MSSQL"].Port)
}
