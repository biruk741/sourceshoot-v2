package migrate

import (
	"backend/data/models"
)

// MigrationModels is a list of all migrations to run
// when the server starts. Each migration should
// be a function that returns an error.
// test
var MigrationModels = []models.Model{
	models.User{},
	models.Industry{},
	models.Skill{},
	models.Business{},
	models.Worker{},
	models.WorkerSkill{},
	models.PrivateParty{},
	models.JobApplication{},
	models.PrivatePartyJobListing{},
	models.BusinessJobListing{},
	models.Questionnaire{},
	models.QuestionnaireQuestion{},
	models.QuestionnaireResponse{},
	models.QuestionnaireOption{},
	models.Shift{},
	models.Review{},
}

// RunMigrations runs all the database migrations
func RunMigrations() error {
	for _, model := range MigrationModels {
		if err := model.RunMigration(); err != nil {
			return err
		}
	}
	return nil
}
