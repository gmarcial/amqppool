package amqppool

import (
	"github.com/streadway/amqp"
)

//ReusableChannel represents a channel amqp that can be reusable
type ReusableChannel struct {
	ID             int           //identification of a reusable channel
	released       bool          //indicates when the channel was released
	channelRelease chan int      //a go channel to notify the pool which the reusable was released
	channel        *amqp.Channel //channel to be reuse
}

//Release release the reusable channel in use back to pool
func (reusableChannel *ReusableChannel) Release() {
	reusableChannel.released = true

	reusableChannel.channelRelease <- reusableChannel.ID
}

//isReleased encapsulate the verification if the reusable channel was released
func (reusableChannel *ReusableChannel) isReleased() error {
	if reusableChannel.released {
		return &UseReleaseChannel{message: "Tried to use a reusable channel that was already released"}
	}

	return nil
}
