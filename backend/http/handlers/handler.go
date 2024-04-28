package handlers

import "backend/services"

type Handler struct {
	services.UserService
	services.SkillsService
	services.IndustryService
	services.LocationService
}
