package models

type Response struct {
	ApiStatus  int         `json:"api_status"`
	ApiMessage string      `json:"api_message"`
	Count      int64       `json:"count"`
	Data       interface{} `json:"data"`
	Token      string      `json:"token"`
}

type Credentials struct {
	Cid     string `json:"client_id"`
	Csecret string `json:"client_secret"`
}

type Status struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
}

type JsonResponseToken struct {
	ApiStatus        int         `gorm:"default:'null'" json:"api_status"`
	ApiMessage       string      `gorm:"default:'null'" json:"api_message"`
	ApiAuthorization string      `gorm:"default:'null'" json:"api_authorization"`
	ApiHttp          int         `gorm:"default:'null'" json:"api_http"`
	Data             interface{} `gorm:"default:'null'" json:"data"`
	Token            string      `json:"token"`
}
