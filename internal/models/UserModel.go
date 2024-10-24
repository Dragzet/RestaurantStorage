package models

type UserModel struct {
	ID                          int    `json:"ID" swaggerignore:"true"`
	Email                       string `json:"Email"`
	Password                    string `json:"Password"`
	CriticalRemainingPercentage int    `json:"CRP"`
}
