package model

type Doctor struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"unique" json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Drug struct {
	DrugID int64  `json:"drug_id" gorm:"primaryKey;autoIncrement"`
	Name   string `json:"name"`
	Price  int64  `json:"price"`
}

type CreatePrescInput struct {
	PreId      int64    `json:"preid"`
	Drugs      []string `json:"drugs"`
	Expiration string   `json:"expiration"`
}

type Opinion struct {
	ID      int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Rating  int    `json:"rating" gorm:"column:rating"`
	Comment string `json:"comment"`
}

type Order struct {
	ID      int64  `json:"order-id" gorm:"primaryKey;autoIncrement"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Address string `json:"address"`
	Items   string `json:"items"`
	Price   string `json:"price"`
}

type User struct {
	ID       int64  `gorm:"primaryKey;autoIncrement"`
	Login    string `gorm:"unique" json:"login"`
	Password string `json:"password"`
	Role     string `json:"role"`
}
