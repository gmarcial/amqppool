package example

/*import (
	"github.com/gmarcial/amqppool"
	"github.com/streadway/amqp"
	"log"
	"os"
	"time"
)

func main() {
	connectionString := os.Getenv("AMQP_CONNECTION")
	maxChannels := 10
	logger := log.New(os.Stdout, "", log.LstdFlags)

	//Action
	pool, err := amqppool.NewPool(connectionString, maxChannels, logger)
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	reusableChannel, err := pool.GetReusableChannel()
	if err != nil {
		panic(err)
	}
	defer reusableChannel.Release()

	publishing := amqp.Publishing{
		ContentType:  "application/json",
		Body:         []byte("Sample"),
		DeliveryMode: amqp.Persistent,
		Timestamp:    time.Now(),
	}

	err = reusableChannel.Publish("exchange", "key", false, true, publishing)
	if err != nil {
		panic(err)
	}
}*/