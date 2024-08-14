package src

type User struct {
	ID         int    `json:"id,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,omitempty"`
	Name       string `json:"name,omitempty"`
	Gender     string `json:"gender,omitempty"`
	Address    string `json:"address,omitempty"`
	IsVerified bool   `json:"is_verified,omitempty"` //bisa diambil dari subscription
}

type Liked struct {
	ID     int `json:"id,omitempty"`
	UserID int `json:"user_id,omitempty"`
}

type SwipedCount struct {
	ID       int `json:"id,omitempty"`
	UserID   int `json:"user_id,omitempty"`
	PersonID int `json:"person_id,omitempty"`
}

type Subscription struct {
	ID     int `json:"id,omitempty"`
	UserID int `json:"user_id,omitempty"`
}
