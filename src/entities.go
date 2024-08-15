package src

import "time"

type Base struct {
	ID        int        `json:"id"`
	CreatedAt *time.Time `json:"created_at"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type User struct {
	Base     Base
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Password string `json:"password,omitempty"`
}

type Profile struct {
	Base      Base
	UserID    int       `json:"user_id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Address   string    `json:"address,omitempty"`
	Photo     string    `json:"photo,omitempty"`
	BirthDate time.Time `json:"birth_date,omitempty"`
	Gender    int       `json:"gender,omitempty"`
	Bio       string    `json:"bio,omitempty"`
}

type Match struct {
	Base     Base
	UserID   string `json:"user_id,omitempty"`
	PersonID string `json:"person_id,omitempty"`
	Status   int    `json:"status,omitempty"`
}

type Subscription struct {
	Base      Base
	UserID    int       `json:"user_id,omitempty"`
	StartDate time.Time `json:"start_date,omitempty"`
	EndDate   time.Time `json:"end_date,omitempty"`
	Status    int       `json:"status,omitempty"`
}

type Payment struct {
	Base           Base
	UserID         int       `json:"user_id,omitempty"`
	SubscriptionID int       `json:"subscription_id,omitempty"`
	Amount         int       `json:"amount,omitempty"`
	PaymentMethod  string    `json:"payment_method,omitempty"`
	Status         int       `json:"status,omitempty"`
	PaymentDate    time.Time `json:"payment_date,omitempty"`
}
