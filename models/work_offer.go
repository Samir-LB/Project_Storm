package models

import (
	"encoding/json"
	"github.com/gofrs/uuid"
	"time"
)

type WorkOffer struct {
	ID          uuid.UUID `json:"id" db:"id"`
	Name        string    `json:"name" db:"name"`
	Description string    `json:"description" db:"description"`
	Address     string    `json:"address" db:"address"`
	Contract    string    `json:"contract" db:"contract"`
	CompanyID   uuid.UUID `json:"company_id" db:"company_id" `
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}

// WorkOffers is not required by pop and may be deleted
type WorkOffers []WorkOffer

// String is not required by pop and may be deleted
func (w WorkOffers) String() string {
	jd, _ := json.Marshal(w)
	return string(jd)
}

func (w WorkOffer) String() string {
	jd, _ := json.Marshal(w)
	return string(jd)
}
