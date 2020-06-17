package amqppool

var (
	ErrAllChannelsInUse = &AllChannelsInUseError{message: "failed in try get a reusable channel, all are in use"}
	ErrUseReleaseChannel = &UseReleaseChannelError{message: "Tried to use a reusable channel that was already released"}
)

//AllChannelsInUseError an error of when is tried to get a reusable channel, but was hit the maximum quantity of pool.
type AllChannelsInUseError struct {
	message string
}

//Error implementing the error interface
func (err *AllChannelsInUseError) Error() string {
	return err.message
}

//UseReleaseChannelError an error of when is tried to use a reusable channel but was released back to pool.
type UseReleaseChannelError struct {
	message string
}

//Error implementing the error interface
func (err *UseReleaseChannelError) Error() string {
	return err.message
}
