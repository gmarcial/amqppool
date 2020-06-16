package amqppool

import "github.com/streadway/amqp"

//Declaration of wrappers of the methods of the channel,
//complementing them with the behavior of the pool.

//Ack wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Ack(tag uint64, multiple bool) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.Ack(tag, multiple)
}

//Reject wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Reject(tag uint64, requeue bool) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.Reject(tag, requeue)
}

//Nack wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Nack(tag uint64, multiple bool, requeue bool) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.Nack(tag, multiple, requeue)
}

//Publish wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Publish(exchange, key string, mandatory, immediate bool, msg amqp.Publishing) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.Publish(exchange, key, mandatory, immediate, msg)
}

//QueueBind wrap to use in reusable channel
func (reusableChannel *ReusableChannel) QueueBind(name, key, exchange string, noWait bool, args amqp.Table) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.QueueBind(name, key, exchange, noWait, args)
}

//ExchangeDeclare wrap to use in reusable channel
func (reusableChannel *ReusableChannel) ExchangeDeclare(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.ExchangeDeclare(name, kind, durable, autoDelete, internal, noWait, args)
}

//Cancel wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Cancel(consumer string, noWait bool) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.Cancel(consumer, noWait)
}

//Confirm wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Confirm(noWait bool) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.Confirm(noWait)
}

//ExchangeBind wrap to use in reusable channel
func (reusableChannel *ReusableChannel) ExchangeBind(destination, key, source string, noWait bool, args amqp.Table) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.ExchangeBind(destination, key, source, noWait, args)
}

//ExchangeDeclarePassive wrap to use in reusable channel
func (reusableChannel *ReusableChannel) ExchangeDeclarePassive(name, kind string, durable, autoDelete, internal, noWait bool, args amqp.Table) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.ExchangeDeclarePassive(name, kind, durable, autoDelete, internal, noWait, args)
}

//ExchangeDelete wrap to use in reusable channel
func (reusableChannel *ReusableChannel) ExchangeDelete(name string, ifUnused, noWait bool) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.ExchangeDelete(name, ifUnused, noWait)
}

//ExchangeUnbind wrap to use in reusable channel
func (reusableChannel *ReusableChannel) ExchangeUnbind(destination, key, source string, noWait bool, args amqp.Table) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.ExchangeUnbind(destination, key, source, noWait, args)
}

//Flow wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Flow(active bool) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.Flow(active)
}

//Qos wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Qos(prefetchCount, prefetchSize int, global bool) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.Qos(prefetchCount, prefetchSize, global)
}

//Recover wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Recover(requeue bool) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.Recover(requeue)
}

//QueueUnbind wrap to use in reusable channel
func (reusableChannel *ReusableChannel) QueueUnbind(name, key, exchange string, args amqp.Table) error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.QueueUnbind(name, key, exchange, args)
}

//Tx wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Tx() error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.Tx()
}

//TxCommit wrap to use in reusable channel
func (reusableChannel *ReusableChannel) TxCommit() error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.TxCommit()
}

//TxRollback wrap to use in reusable channel
func (reusableChannel *ReusableChannel) TxRollback() error {
	if err := reusableChannel.isReleased(); err != nil {
		return err
	}

	return reusableChannel.channel.TxRollback()
}

//Get wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Get(queue string, autoAck bool) (msg amqp.Delivery, ok bool, err error) {
	if err := reusableChannel.isReleased(); err != nil {
		return amqp.Delivery{}, false, err
	}

	return reusableChannel.channel.Get(queue, autoAck)
}

//Consume wrap to use in reusable channel
func (reusableChannel *ReusableChannel) Consume(queue, consumer string, autoAck, exclusive, noLocal, noWait bool, args amqp.Table) (<-chan amqp.Delivery, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return nil, err
	}

	return reusableChannel.channel.Consume(queue, consumer, autoAck, exclusive, noLocal, noWait, args)
}

//QueueInspect wrap to use in reusable channel
func (reusableChannel *ReusableChannel) QueueInspect(name string) (amqp.Queue, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return amqp.Queue{}, err
	}

	return reusableChannel.channel.QueueInspect(name)
}

//NotifyClose wrap to use in reusable channel
func (reusableChannel *ReusableChannel) NotifyClose(c chan *amqp.Error) (chan *amqp.Error, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return nil, err
	}

	return reusableChannel.channel.NotifyClose(c), nil
}

//NotifyCancel wrap to use in reusable channel
func (reusableChannel *ReusableChannel) NotifyCancel(c chan string) (chan string, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return nil, err
	}

	return reusableChannel.channel.NotifyCancel(c), nil
}

//NotifyConfirm wrap to use in reusable channel
func (reusableChannel *ReusableChannel) NotifyConfirm(ack, nack chan uint64) (chan uint64, chan uint64, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return nil, nil, err
	}

	ack, nack = reusableChannel.channel.NotifyConfirm(ack, nack)

	return ack, nack, nil
}

//QueueDeclare wrap to use in reusable channel
func (reusableChannel *ReusableChannel) QueueDeclare(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return amqp.Queue{}, err
	}

	return reusableChannel.channel.QueueDeclare(name, durable, autoDelete, exclusive, noWait, args)
}

//NotifyFlow wrap to use in reusable channel
func (reusableChannel *ReusableChannel) NotifyFlow(c chan bool) (chan bool, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return nil, err
	}

	return reusableChannel.channel.NotifyFlow(c), nil
}

//NotifyPublish wrap to use in reusable channel
func (reusableChannel *ReusableChannel) NotifyPublish(confirm chan amqp.Confirmation) (chan amqp.Confirmation, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return nil, err
	}

	return reusableChannel.channel.NotifyPublish(confirm), nil
}

//NotifyReturn wrap to use in reusable channel
func (reusableChannel *ReusableChannel) NotifyReturn(c chan amqp.Return) (chan amqp.Return, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return nil, err
	}

	return reusableChannel.channel.NotifyReturn(c), nil
}

//QueueDeclarePassive wrap to use in reusable channel
func (reusableChannel *ReusableChannel) QueueDeclarePassive(name string, durable, autoDelete, exclusive, noWait bool, args amqp.Table) (amqp.Queue, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return amqp.Queue{}, err
	}

	return reusableChannel.channel.QueueDeclarePassive(name, durable, autoDelete, exclusive, noWait, args)
}

//QueueDelete wrap to use in reusable channel
func (reusableChannel *ReusableChannel) QueueDelete(name string, ifUnused, ifEmpty, noWait bool) (int, error) {
	if err := reusableChannel.isReleased(); err != nil {
		return 0, err
	}

	return reusableChannel.channel.QueueDelete(name, ifUnused, ifEmpty, noWait)
}
