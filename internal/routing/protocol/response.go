package protocol

type Response interface {
	Type() MessageType
}

type ProduceResponse struct {
	Offset int64
	Err    string
}

func (r *ProduceResponse) Type() MessageType {
	return MsgProduce
}
