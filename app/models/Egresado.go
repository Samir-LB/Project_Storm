package models

import (
	"github.com/gobuffalo/validate/v3"
	"github.com/gobuffalo/validate/v3/validators"
)

type Egresado struct {
	ID          string `json:"id" db:"id"`
	Nombre      string `json:"nombre" db:"nombre"`
	Correo      string `json:"correo" db:"correo"`
	Clave       string `json:"clave" db:"clave"`
	Descripcion string `json:"descripcion" db:"descripcion"`
	Direccion   string `json:"direccion" db:"direccion"`
	Contrato    string `json:"contrato" db:"contrato"`
}

func (e *Egresado) Validate() (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: e.Nombre, Name: "Nombre"},
		&validators.StringLengthInRange{Field: e.Nombre, Name: "Nombre", Min: 1, Max: 255},
		&validators.StringIsPresent{Field: e.Descripcion, Name: "Descripcion"},
		&validators.StringLengthInRange{Field: e.Descripcion, Name: "Descripcion", Min: 0, Max: 100},
		&validators.StringIsPresent{Field: e.Direccion, Name: "Direccion"},
		&validators.StringLengthInRange{Field: e.Direccion, Name: "Direccion", Min: 1, Max: 255},
		&validators.StringIsPresent{Field: e.Contrato, Name: "Contrato"},
		&validators.StringIsPresent{Field: e.Correo, Name: "Correo"},
		&validators.EmailIsPresent{Field: e.Correo, Name: "Correo"},
		&validators.StringIsPresent{Field: e.Clave, Name: "Clave"},
		&validators.StringLengthInRange{Field: e.Clave, Name: "Clave", Min: 6, Max: 30},
	), nil
}
