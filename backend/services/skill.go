package services

import (
	"fmt"

	"backend/data/models"
	"backend/data/repo"
	serviceTypes "backend/services/types"
)

type SkillsService interface {
	GetSkills() ([]serviceTypes.Skill, error)
}

type SkillInstance struct {
	repo.SkillRepo
}

func NewSkillService(skillRepo repo.SkillRepo) *SkillInstance {
	s := SkillInstance{skillRepo}
	return &s
}

func (s *SkillInstance) GetSkills() ([]serviceTypes.Skill, error) {
	skills, err := s.SkillRepo.GetSkillsFromDB()
	if err != nil {
		return nil, err
	}
	return ConvertGormSkillToService(skills)
}

func ConvertGormSkillToService(s []models.Skill) ([]serviceTypes.Skill, error) {
	skills := make([]serviceTypes.Skill, len(s))
	for i, skill := range s {
		empty := models.Industry{}
		if skill.Industry == empty {
			return []serviceTypes.Skill{}, fmt.Errorf("ConvertGormSkillToService: industry is empty")
		}
		_industry := serviceTypes.Industry{
			IndustryID:   skill.Industry.IndustryID,
			IndustryName: skill.Industry.IndustryName,
			Description:  skill.Industry.Description,
		}
		skills[i] = serviceTypes.Skill{
			ID:          skill.Model.ID,
			Name:        skill.SkillName,
			Description: skill.Description,
			Industry:    _industry,
		}
	}
	return skills, nil // todo: return err if occurs
}
