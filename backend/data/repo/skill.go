package repo

import (
	"gorm.io/gorm"

	"backend/data"
	"backend/data/models"
)

type SkillRepo interface {
	CreateSkill(skill models.Skill) (uint, error)
	GetSkills() ([]models.Skill, error)
}

type SkillRepoInstance struct {
	db *gorm.DB
}

func NewSkillRepo() SkillRepoInstance {
	return SkillRepoInstance{db: data.DB}
}

func (r SkillRepoInstance) CreateSkill(skill models.Skill) (uint, error) {
	db := data.DB
	query := db.Create(&skill)
	if err := query.Error; err != nil {
		return 0, err
	}

	return skill.ID, nil
}

func (r SkillRepoInstance) GetSkills() ([]models.Skill, error) {
	skills := []models.Skill{}
	if err := r.db.Preload("Industry").Find(&skills).Error; err != nil {
		return []models.Skill{}, err
	}
	return skills, nil
}
