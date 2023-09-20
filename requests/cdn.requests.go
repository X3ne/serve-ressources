package requests

import (
	validation "github.com/go-ozzo/ozzo-validation"
)

type RessourceBimg struct {
	Q	int	`query:"q" validate:"optional" example:"?q=50" description:"Quality of the image" minimum:"1" maximum:"100"`
	A	bool	`query:"a" validate:"optional" example:"?a=true" description:"Animated image" enum:"true,false"`
	C	bool	`query:"c" validate:"optional" example:"?c=true" description:"Crop image" enum:"true,false"`
	W int		`query:"w" validate:"optional" example:"?w=100" description:"Width of the image" minimum:"1"`
	H int		`query:"h" validate:"optional" example:"?h=100" description:"Height of the image" minimum:"1"`
}

func (rb RessourceBimg) Validate() error {
	return validation.ValidateStruct(&rb,
		validation.Field(&rb.Q, validation.Min(1), validation.Max(100)),
		validation.Field(&rb.A),
		validation.Field(&rb.C),
		validation.Field(&rb.W, validation.Min(1)),
		validation.Field(&rb.H, validation.Min(1)),
	)
}
