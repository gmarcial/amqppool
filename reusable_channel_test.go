package amqppool

import (
	"log"
	"os"
	"testing"
)

func TestShouldReleaseAReusableChannel(t *testing.T){
	//Arrange
	connectionString := os.Getenv("AMQP_CONNECTION")
	maxChannels := 10
	logger := log.New(os.Stdout, "", log.LstdFlags)

	pool, _ := NewPool(connectionString, maxChannels, logger)
	defer pool.Close()

	reusableChannel, _ := pool.GetReusableChannel()

	//Action
	reusableChannel.Release()

	//Assert
	if reusableChannel.released != true {
		t.Error("The reusable channel don't was released")
	}
}
