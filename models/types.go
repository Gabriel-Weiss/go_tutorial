package models

type Product struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Image       string  `json:"image"`
	Rating      Rating  `json:"rating"`
}

type Rating struct {
	Rate  float64 `json:"rate"`
	Count int     `json:"count"`
}

type CartProduct struct {
	ProductID int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type Cart struct {
	ID       int           `json:"id"`
	UserID   int           `json:"userId"`
	Date     string        `json:"date"` // Adjust to `time.Time` if using proper date format
	Products []CartProduct `json:"products"`
}

type User struct {
	ID       int         `json:"id"`
	Email    string      `json:"email"`
	Username string      `json:"username"`
	Password string      `json:"password"`
	Name     FullName    `json:"name"`
	Address  UserAddress `json:"address"`
	Phone    string      `json:"phone"`
}

type FullName struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type UserAddress struct {
	City        string      `json:"city"`
	Street      string      `json:"street"`
	Number      int         `json:"number"`
	ZipCode     string      `json:"zipcode"`
	GeoLocation GeoLocation `json:"geolocation"`
}

type GeoLocation struct {
	Lat  string `json:"lat"`
	Long string `json:"long"`
}
