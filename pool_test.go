package amqppool

import (
	"errors"
	"log"
	"os"
	"testing"
)

func TestShouldCreateANewAmqpPool(t *testing.T) {
	//Arrange
	connectionString := os.Getenv("AMQP_CONNECTION")
	maxChannels := 10
	logger := log.New(os.Stdout, "", log.LstdFlags)

	//Action
	pool, err := NewPool(connectionString, maxChannels, logger)
	defer pool.Close()

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to create a new pool: %v", err.Error())
	}

	if pool.maxChannels != maxChannels {
		t.Errorf("The maximum quantity of channels is inconsistent: Expected %v and found %v",
			maxChannels, pool.maxChannels)
	}

	lenChannelsReleased := len(pool.channelsReleased)
	if lenChannelsReleased != maxChannels {
		t.Errorf("The quantity of channels opend is inconsistent: Expected %v and found %v",
			maxChannels, lenChannelsReleased)
	}
}

func TestShouldReturnAErrorWhenCreateANewAmqpPoolWhenTheConnectionFail(t *testing.T) {
	//Arrange
	connectionString := os.Getenv("AMQP_CONNECTION_WRONG")
	maxChannels := 10
	logger := log.New(os.Stdout, "", log.LstdFlags)

	//Action
	pool, err := NewPool(connectionString, maxChannels, logger)

	//Assert
	if pool != nil && err == nil {
		t.Error("Don't occurred an error to create a new pool: Was created with a fail connection")
	}
}

func TestShouldCloseAAmqpPool(t *testing.T) {
	//Arrange
	connectionString := os.Getenv("AMQP_CONNECTION")
	maxChannels := 10
	logger := log.New(os.Stdout, "", log.LstdFlags)

	pool, _ := NewPool(connectionString, maxChannels, logger)

	//Action
	err := pool.Close()

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to close the pool: %v", err.Error())
	}

}

func TestShouldCloseAReusableChannelByTheYourId(t *testing.T) {
	//Arrange
	connectionString := os.Getenv("AMQP_CONNECTION")
	maxChannels := 10
	logger := log.New(os.Stdout, "", log.LstdFlags)

	pool, _ := NewPool(connectionString, maxChannels, logger)
	defer pool.Close()

	//Action
	err := pool.CloseReusableChannel(1)

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to close a reusable channel: %v", err.Error())
	}
}

func TestShouldGetAReusableChannelWhenThePoolHaveAnyChannelReleased(t *testing.T) {
	//Arrange
	connectionString := os.Getenv("AMQP_CONNECTION")
	maxChannels := 10
	logger := log.New(os.Stdout, "", log.LstdFlags)

	pool, _ := NewPool(connectionString, maxChannels, logger)
	defer pool.Close()

	//Action
	reusableChannel, err := pool.GetReusableChannel()
	defer reusableChannel.Release()

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to get a reusable channel: %v", err.Error())
	}

	if reusableChannel.released == true {
		t.Errorf("The reusable channel was obtained how released")
	}
}

func TestShouldGetAReusableChannelWhenThePoolDontHaveAnyChannelReleasedButIsNotAtTheLimitOfUse(t *testing.T) {
	//Arrange
	connectionString := os.Getenv("AMQP_CONNECTION")
	maxChannels := 1
	logger := log.New(os.Stdout, "", log.LstdFlags)

	pool, _ := NewPool(connectionString, maxChannels, logger)
	defer pool.Close()
	_ = pool.CloseReusableChannel(1)

	//Action
	reusableChannel, err := pool.GetReusableChannel()
	defer reusableChannel.Release()

	//Assert
	if err != nil {
		t.Errorf("Occurred a error to get a reusable channel: %v", err.Error())
	}

	if reusableChannel.released == true {
		t.Errorf("The reusable channel was obtained how released")
	}
}

func TestShouldReturnErrorInGetAReusableChannelWhenTheAllChannelAreInUse(t *testing.T) {
	//Arrange
	connectionString := os.Getenv("AMQP_CONNECTION")
	maxChannels := 1
	logger := log.New(os.Stdout, "", log.LstdFlags)

	pool, _ := NewPool(connectionString, maxChannels, logger)
	defer pool.Close()

	reusableChannel, err := pool.GetReusableChannel()
	defer reusableChannel.Release()

	//Action
	reusableChannel, err = pool.GetReusableChannel()

	//Assert
	if err == nil {
		t.Error("Don't occurred a error to get a reusable channel when all are in use.")
	}

	if reusableChannel != nil {
		t.Error("The reusable channel was obtained same when all are in use")
	}

	if !errors.Is(err, ErrAllChannelsInUse){
		t.Error("The type of error returned is different of expected")
	}
}