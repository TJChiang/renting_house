package house591

type HouseStructure struct {
	BluekaiData map[string]interface{} `json:"bluekai_data"`
	Data        MainData               `json:"data"`
	Status      int                    `json:"status"`
}

type MainData struct {
	TopData []map[string]interface{} `json:"topData"`
	Data    []map[string]interface{} `json:"data"`
}
