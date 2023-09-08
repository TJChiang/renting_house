package house591

type HouseStructure struct {
	BluekaiData map[string]interface{} `json:"bluekai_data"`
	Data        MainData               `json:"data"`
	Status      int                    `json:"status"`
}

type MainData struct {
	TopData []MainDataElement `json:"topData"`
	Data    []MainDataElement `json:"data"`
}

type MainDataElement struct {
	PostId      int      `json:"post_id"`
	Title       string   `json:"title"`
	KindName    string   `json:"kind_name"`
	Type        string   `json:"type"`
	Location    string   `json:"location"`
	SectionName string   `json:"section_name"`
	StreetName  string   `json:"street_name"`
	Floor       string   `json:"floor_str"`
	Room        string   `json:"room_str"`
	Price       string   `json:"price"`
	PriceUnit   string   `json:"price_unit"`
	RoleName    string   `json:"role_name"`
	Contact     string   `json:"contact"`
	ImageList   []string `json:"photo_list"`
	RefreshTime string   `json:"refresh_time"`
	Community   string   `json:"community"`
	Area        string   `json:"area"`
	Cid         string   `json:"cid"`
	Tag         []Tag    `json:"rent_tag"`
	IsHurry     int      `json:"hurry"`
	IsCombine   int      `json:"is_combine"`
	IsSocial    int      `json:"is_social"`
	IsVip       int      `json:"is_vip"`
	Preferred   int      `json:"preferred"`
	Hit         int      `json:"yesterday_hit"`
}

type Tag struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
