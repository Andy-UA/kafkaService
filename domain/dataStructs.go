package domain

type Flatters struct {
	InputTopic         string             `json:"inputTopic"`
	DestinationTopic   string             `json:"destinationTopic"`
	Graph              Graph              `json:"graph"`
	DestinationMessage DestinationMessage `json:"destinationMessage"`
}

type Metric struct {
	UsedSpaceBytes  interface{} `json:"usedSpaceBytes"`
	TotalSpaceBytes interface{} `json:"totalSpaceBytes"`
}

type BasicInfo struct {
	Name      string 	  `json:"name"`
	DriveType interface{} `json:"driveType"`
}

type InputPartition struct {
	BasicInfo
	Metric `json:"metric"`
}

type OutputPartition struct {
	BasicInfo
	Metric
}

type Message struct {
	CreateAtTimeUTC string      	 `json:"createAtTimeUTC"`
	Partitions      []InputPartition `json:"partitions"`
}

type Graph struct {
	Message Message `json:"Message"`
}

type Data struct {
	OutputPartition
	CreateAtTimeUTC string `json:"createAtTimeUTC"`
}

type DestinationMessage struct {
	Data Data `json:"data"`
}

type InputMessage struct {
	Action string
	Message Message
}

type OutputMessage struct {
	PayloadData Data `json:"payloadData"`
}
