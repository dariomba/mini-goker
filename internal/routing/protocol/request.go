package protocol

type MessageType uint8

type Request interface {
	Type() MessageType
}

type ProduceRequest struct {
	Topic     string
	Partition int
	Payload   []byte
}

func (r *ProduceRequest) Type() MessageType {
	return MsgProduce
}
