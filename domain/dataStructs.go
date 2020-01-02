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

type Partitions struct {
	Name      string `json:"name"`
	DriveType interface{} `json:"driveType"`
	Metric    Metric `json:"metric"`
}

type Message struct {
	CreateAtTimeUTC string       `json:"createAtTimeUTC"`
	Partitions      []Partitions `json:"partitions"`
}

type Graph struct {
	Message Message `json:"Message"`
}

type Data struct {
	Name            string `json:"name"`
	DriveType       interface{} `json:"driveType"`
	UsedSpaceBytes  interface{} `json:"usedSpaceBytes"`
	TotalSpaceBytes interface{} `json:"totalSpaceBytes"`
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