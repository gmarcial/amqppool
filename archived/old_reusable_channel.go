package archived

/*import (
"github.com/streadway/amqp"
"log"
)

//ReusableChannel represents a channel amqp that can be reusable
type ReusableChannel struct {
	ID             int
	released       bool
	channelRelease chan int
	channel        *amqp.Channel
	reference      **amqp.Channel
}

//UseChannel get the reusable channel
func (reusableChannel *ReusableChannel) UseChannel(channel **amqp.Channel) {
	if reusableChannel.released {
		log.Panicf("Tried to use a reusable channel that was already released")
	}

	log.Printf("Go the use the reusable channel %v", reusableChannel.ID)
	reusableChannel.reference = channel
	*channel = reusableChannel.channel
}

//Release release the reusable channel in use
func (reusableChannel *ReusableChannel) Release() {
	reusableChannel.released = true
	*reusableChannel.reference = nil

	log.Printf("The reusable channel %v was released with success", reusableChannel.ID)
	reusableChannel.channelRelease <- reusableChannel.ID
}*/
