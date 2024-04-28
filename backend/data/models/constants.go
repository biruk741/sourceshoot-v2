package models

type PayType string

const (
	Hourly  PayType = "hourly"
	Weekly  PayType = "weekly"
	Monthly PayType = "monthly"
	Yearly  PayType = "yearly"
)
