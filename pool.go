package amqppool

import (
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

//Pool represents a connection and manage the pool of reusable channels
type Pool struct {
	connection                  *amqp.Connection         //the connection amqp
	maxChannels                 int                      //the maximum quantity of channels of pool
	channelRelease              chan int                 //a go channel to listen when a reusable channel was released
	connectionCloseNotification chan *amqp.Error         //a go channel to listen when the connection amqp was closed
	channelsInUse               map[int]*ReusableChannel //in use channels store
	channelsReleased            map[int]*ReusableChannel //released channels store
}

//NewPool create a new Pool
func NewPool(connectionString string, maxChannels int, logger *log.Logger) (*Pool, error) {
	connection, err := connect(connectionString)
	if err != nil {
		return nil, err
	}
	reusableChannels := make(map[int]*ReusableChannel, 0)
	channelRelease := make(chan int)

	for id := 1; id <= maxChannels; id++ {
		reusableChannel, err := newReusableChannel(id, connection, channelRelease)
		if err != nil {
			return nil, err
		}

		reusableChannels[id] = reusableChannel
	}

	connectionCloseNotification := make(chan *amqp.Error)
	connectionCloseNotification = connection.NotifyClose(connectionCloseNotification)

	pool := &Pool{
		connection:                  connection,
		maxChannels:                 maxChannels,
		channelRelease:              channelRelease,
		connectionCloseNotification: connectionCloseNotification,
		channelsReleased:            reusableChannels,
		channelsInUse:               make(map[int]*ReusableChannel, 0),
	}

	go listenWhenConnectionClose(connectionString, pool, logger)
	go listenWhenChannelRelease(pool, logger)

	return pool, nil
}

//connect establish the connection with the broker amqp
func connect(connectionString string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(connectionString)
	if err != nil {
		return nil, err
	}

	return connection, nil
}

//Close close the connection with the broker amqp
func (pool *Pool) Close() error {
	connection := pool.connection

	if err := connection.Close(); err != nil {
		errMsg := fmt.Sprintf("Occurred an error to try close the connection with the amqp broker: %v", err.Error())
		return fmt.Errorf(errMsg)
	}

	close(pool.channelRelease)

	return nil
}

//GetReusableChannel get a reusable channel of the pool to use
func (pool *Pool) GetReusableChannel() (*ReusableChannel, error) {
	channelsReleased := pool.channelsReleased
	lenChannelsReleased := len(channelsReleased)
	lenChannelsInUse := len(pool.channelsInUse)

	var reusableChannel *ReusableChannel
	if lenChannelsReleased > 0 {
		for _, channel := range channelsReleased {
			reusableChannel = channel
			pool.addReusableChannelToUse(reusableChannel)
			break
		}
	} else if (lenChannelsReleased == 0) && (lenChannelsInUse < pool.maxChannels) {
		newReusableChannel, err := pool.newReusableChannelToUso()
		reusableChannel = newReusableChannel
		if err != nil {
			return nil, err
		}

		pool.addReusableChannelToUse(reusableChannel)
	} else {
		return nil, ErrAllChannelsInUse
	}

	return reusableChannel, nil
}

//newReusableChannel create a new reusable channel released
func newReusableChannel(id int, connection *amqp.Connection, channelRelease chan int) (*ReusableChannel, error) {
	channel, err := connection.Channel()
	if err != nil {
		return nil, err
	}

	reusableChannel := &ReusableChannel{
		ID:             id,
		released:       true,
		channel:        channel,
		channelRelease: channelRelease,
	}

	return reusableChannel, nil
}

//newReusableChannelToUso create a new reusable channel for now use
func (pool *Pool) newReusableChannelToUso() (*ReusableChannel, error) {
	ID := (len(pool.channelsReleased) + len(pool.channelsInUse)) + 1
	return newReusableChannel(ID, pool.connection, pool.channelRelease)
}

//addReusableChannelToUse add a reusable channel for now use
func (pool *Pool) addReusableChannelToUse(reusableChannel *ReusableChannel) {
	reusableChannel.released = false
	pool.channelsInUse[reusableChannel.ID] = reusableChannel
	delete(pool.channelsReleased, reusableChannel.ID)
}

//Close close a channel and remove of pool
func (pool *Pool) CloseReusableChannel(id int) error {
	reusableChannel, exist := pool.channelsReleased[id]
	if !exist {
		errMsg := fmt.Sprintf("don't was found the reusable channel with the id %v in channels released", id)
		return errors.New(errMsg)
	}

	channel := reusableChannel.channel
	if err := channel.Close(); err != nil {
		errMsg := fmt.Sprintf("Occurred an error to try close the channel of id %v: %v", id, err.Error())
		return errors.New(errMsg)
	}
	delete(pool.channelsReleased, id)

	return nil
}

//listenWhenChannelRelease stay listen when the channels that are released
func listenWhenChannelRelease(pool *Pool, logger *log.Logger) {
	logger.Println("Start listening when reusable channels are released")

	for reusableChannelID := range pool.channelRelease {
		reusableChannel := pool.channelsInUse[reusableChannelID]
		delete(pool.channelsInUse, reusableChannel.ID)
		pool.channelsReleased[reusableChannel.ID] = reusableChannel

		logger.Printf("Reusable channel was released %v", reusableChannelID)
	}

	logger.Println("Stop listening when reusable channels are released")
}

//listenWhenConnectionClose stay listen when the connection amqp close and try to reconnect
func listenWhenConnectionClose(connectionString string, pool *Pool, logger *log.Logger) {
	logger.Println("Start listening when the connection amqp close")

	for err := range pool.connectionCloseNotification {
		logger.Printf("Connection with the broker closed in server %v: %v, %v, try reconnect", err.Server, err.Code, err.Reason)
		connection, err := connect(connectionString)
		if err != nil {
			logger.Panic(err.Error())
		}

		pool.connection = connection
	}

	logger.Println("Stop of listening when the connection amqp close")
}
