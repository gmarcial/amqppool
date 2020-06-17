package amqppool

import (
	"errors"
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

func TestShouldReturnErrorWhenTryingUseAReleasedChannel(t *testing.T){
	//Arrange
	connectionString := os.Getenv("AMQP_CONNECTION")
	maxChannels := 10
	logger := log.New(os.Stdout, "", log.LstdFlags)

	pool, _ := NewPool(connectionString, maxChannels, logger)
	defer pool.Close()

	reusableChannel, _ := pool.GetReusableChannel()
	reusableChannel.Release()

	//Action
	_, _, err := reusableChannel.Get("queue", false)

	//Assert
	if err == nil {
		t.Error("Don't returned a error in to use a channel released")
	}

	if !errors.Is(err, ErrUseReleaseChannel){
		t.Error("The type of error returned is different of expected")
	}

	if reusableChannel.released != true {
		t.Error("The reusable channel don't was released")
	}
}
