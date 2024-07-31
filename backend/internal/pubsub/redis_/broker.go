package gqlgen

// Broker ...
type Broker interface {
	// Subscribe subscribes messages by channelName. It will handle payload, then send handled message to subscription channel.
	Subscribe(channelName string, channel interface{}, payloadHandler PayloadHandler) (Unsubscriber, error)
	// Receive blocks to receive data from Publisher
	Receive() error
}
