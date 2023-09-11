package docker

import (
	"context"
	"github.com/sirupsen/logrus"
	"github.com/testcontainers/testcontainers-go"
)

/**
 * https://golang.testcontainers.org/quickstart/
 */
func CreateContainer(imgName, port string) {
	ctx := context.Background()
	req := testcontainers.ContainerRequest{
		Image:        imgName,
		ExposedPorts: []string{port},
		//WaitingFor: wait.ForLog(""),
	}
	_, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	})
	if nil != err {
		logrus.Error(err)
	}
	//defer func() {
	//	if err := redisC.Terminate(ctx); err != nil {
	//		t.Fatalf("failed to terminate container: %s", err.Error())
	//	}
	//}()
}
