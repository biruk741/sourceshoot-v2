package services

import (
	"backend/data/models"
	"backend/data/repo"
	serviceTypes "backend/services/types"
)

type SkillsService interface {
	GetSkills() ([]models.Skill, error)
}

type SkillInstance struct {
	repo.SkillRepo
}

func NewSkillService(skillRepo repo.SkillRepo) *SkillInstance {
	s := SkillInstance{skillRepo}
	return &s
}

func (s *SkillInstance) GetSkills() ([]models.Skill, error) {
	return s.SkillRepo.GetSkills()
}

func ConvertGormSkillToService(s []models.Skill) ([]serviceTypes.Skill, error) {
	skills := make([]serviceTypes.Skill, len(s))
	for i, skill := range s {
		// empty := models.Industry{}
		// if skill.Industry == empty {
		// 	return []serviceTypes.Skill{}, fmt.Errorf("ConvertGormSkillToService: industry is empty")
		// }
		_industry := serviceTypes.Industry{
			IndustryID:   skill.Industry.ID,
			IndustryName: skill.Industry.Name,
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
