package docker

type DockerContainer struct {
	data map[string]DockerParam
}

type DockerParam struct {
	Name string            `yaml:"name"`
	Port string            `yaml:"port"`
	Env  map[string]string `yaml:"env"`
}
