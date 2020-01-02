package domain

const (
	KafkaServerAddress string = "localhost:9092"
	GroupID string = "messageParserGroup"
	MinByteMessage int = 1e3 //1KB
	MaxByteMessage int = 1e6 //1MB
)
