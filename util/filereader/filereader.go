package filereader

import (
	"github.com/adel5210/goContainer/model/docker"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
)

func ReadDockerYML(filePath string) map[string]docker.DockerParam {
	yFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	data := make(map[string]docker.DockerParam)

	err2 := yaml.Unmarshal(yFile, &data)
	if err2 != nil {
		log.Fatal(err2)
	}
	return data
}
