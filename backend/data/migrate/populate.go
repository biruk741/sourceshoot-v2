package migrate

import (
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"gorm.io/gorm"

	"backend/data"
	"backend/data/models"
	"backend/data/repo"
)

type PopulateFunc func(db *gorm.DB, numData *int, args ...interface{}) error
type PopulateFuncAndAmount struct {
	Func   PopulateFunc
	Amount int
	args   []interface{}
}

const (
	// MOCK_SEED is the seed used for the mock data.
	// This is used to ensure that the same data is generated each time.
	MOCK_SEED = 1233
)

// Populate the database with mock data. The order of the functions in this array matters.
var PopulationsAndAmounts = []PopulateFuncAndAmount{
	{populateSkills, 15, nil},
	{populateWorkers, 10, nil},
	{populateUsers, 10, nil},
	{populateIndustries, 20, nil},
	{Func: populateBusinessJobListings, Amount: 100},
	{Func: populateBusinesses, Amount: 20},
	{Func: PopulateReviews, Amount: 10},
	{Func: PopulateShiftsForWorkerAndBusiness, Amount: 20, args: []interface{}{1, 3}},
}

func populateUsers(db *gorm.DB, numUsers *int, args ...interface{}) error {
	gofakeit.Seed(MOCK_SEED)

	if numUsers == nil {
		defaultNum := 10
		numUsers = &defaultNum
	}

	profileURL := gofakeit.ImageURL(100, 100)
	loginDate := gofakeit.Date()

	for i := 0; i < *numUsers; i++ {
		user := models.User{
			FirebaseID:     gofakeit.UUID(),
			Email:          gofakeit.Email(),
			PhoneNumber:    gofakeit.Phone(),
			LastLogin:      &loginDate,
			UserType:       models.UserType(gofakeit.Number(0, 2)),
			ProfilePicture: &profileURL,
		}

		// Create a new User record in the database
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	return nil
}

func populateSkills(db *gorm.DB, numSkills *int, args ...interface{}) error {
	gofakeit.Seed(MOCK_SEED)

	if numSkills == nil {
		defaultNum := 10
		numSkills = &defaultNum
	}

	for i := 0; i < *numSkills; i++ {
		skill := models.Skill{
			SkillName:   gofakeit.JobTitle(),
			Description: gofakeit.HipsterSentence(5),
			IndustryID:  uint(i + 1),
		}

		// Create a new User record in the database
		if err := db.Create(&skill).Error; err != nil {
			return err
		}
	}

	return nil
}

func populateWorkers(db *gorm.DB, numWorkers *int, _ ...interface{}) error {
	gofakeit.Seed(MOCK_SEED)

	if numWorkers == nil {
		defaultNum := 10
		numWorkers = &defaultNum
	}

	for i := 0; i < *numWorkers; i++ {
		worker := models.Worker{
			UserID:    uint(i),
			FirstName: gofakeit.FirstName(),
			LastName:  gofakeit.LastName(),
			DOB:       gofakeit.Date().UTC(),
			ResumeURL: gofakeit.URL(),
			Bio:       gofakeit.Paragraph(3, 5, 50, "\n"),
			Age:       uint(gofakeit.Number(18, 65)),
			Rating:    gofakeit.Float32Range(1.0, 5.0),
			WorkType:  models.WorkType(models.FULL_TIME),
			MinPay:    gofakeit.Float32Range(10.0, 50.0),
			SSN:       gofakeit.SSN(),
			// IndustryID: uint(gofakeit.IntRange(0, 10)),
		}

		// Create a new Worker record in the database
		if err := db.Create(&worker).Error; err != nil {
			return err
		}
	}

	return nil
}

func populateIndustries(db *gorm.DB, numIndustries *int, args ...interface{}) error {
	gofakeit.Seed(MOCK_SEED)

	if numIndustries == nil {
		defaultNum := 20 // Increase the default number to 20
		numIndustries = &defaultNum
	}

	industries := []string{
		"Engineering",
		"Education",
		"Tech",
		"Healthcare",
		"Finance",
		"Entertainment",
		"Retail",
		"Hospitality",
		"Automotive",
		"Aerospace",
		"Construction",
		"Media",
		"Energy",
		"Telecommunications",
		"Biotechnology",
		"Fashion",
		"Sports",
		"Real Estate",
		"Transportation",
	}

	for i := 0; i < len(industries); i++ {
		industry := models.Industry{
			IndustryName: industries[i],
			Description:  gofakeit.Sentence(5),
		}

		// Create a new Industry record in the database
		if err := db.Create(&industry).Error; err != nil {
			return err
		}
	}

	return nil
}

func populateBusinesses(db *gorm.DB, numBusinesses *int, args ...interface{}) error {
	gofakeit.Seed(MOCK_SEED)

	if numBusinesses == nil {
		defaultNum := 20 // Increase the default number to 20
		numBusinesses = &defaultNum
	}

	for i := 0; i < *numBusinesses; i++ {
		b := models.Business{
			UserID:              uint(gofakeit.IntRange(2, 4)),
			BusinessName:        gofakeit.Company(),
			BusinessPhoneNumber: gofakeit.Phone(),
			EIN:                 gofakeit.SSN(),
			LastLogin:           gofakeit.DateRange(time.Now().AddDate(0, 0, -30), time.Now()),
			Rating:              gofakeit.Float32Range(1.0, 5.0),
			NumEmployees:        uint(gofakeit.IntRange(1, 100)),
			BusinessAddress: models.Address{
				Street:  gofakeit.Street(),
				Street2: gofakeit.Street(),
				City:    gofakeit.City(),
				State:   gofakeit.State(),
				ZipCode: gofakeit.Zip(),
			},
			// IndustryID:  uint(gofakeit.IntRange(0, 10)),
			Description: gofakeit.SentenceSimple(),
		}

		// Create a new Industry record in the database
		if err := db.Create(&b).Error; err != nil {
			return err
		}
	}

	return nil
}

func populateBusinessJobListings(db *gorm.DB, numListings *int, args ...interface{}) error {
	gofakeit.Seed(MOCK_SEED)

	if numListings == nil {
		defaultNum := 20 // Increase the default number to 20
		numListings = &defaultNum
	}

	for i := 0; i < *numListings; i++ {
		desc := gofakeit.LoremIpsumSentence(10)
		lat, err := gofakeit.LatitudeInRange(44.47, 45.42)
		if err != nil {
			return err
		}
		long, err := gofakeit.LongitudeInRange(-94.01, -92.73)
		if err != nil {
			return err
		}
		address := gofakeit.Address()
		listing := models.BusinessJobListing{
			BusinessID:         uint(i),
			ListingTitle:       gofakeit.JobTitle(),
			ListingDescription: &desc,
			PayType:            models.Hourly,
			Location: data.GeoLocation{
				Latitude:  lat,
				Longitude: long,
			},
			Address: data.Address{
				Street:  address.Street,
				Street2: address.Street,
				City:    address.City,
				State:   address.State,
				ZipCode: address.Zip,
			},
			PayAmount:    uint(gofakeit.Float32Range(10.0, 50.0)),
			DeadlineDate: gofakeit.DateRange(time.Now(), time.Now().AddDate(0, 0, 30)),
		}

		// Create a new Industry record in the database
		if err := db.Create(&listing).Error; err != nil {
			return err
		}
	}

	return nil
}

// GenerateFakeShifts creates and inserts fake shifts in the database
func PopulateShiftsForWorkerAndBusiness(db *gorm.DB, numShifts *int, args ...interface{}) error {
	const (
		defaultWorkerID   int = 0
		defaultBusinessID int = 0
	)

	args = args[0].([]interface{})

	workerID := defaultWorkerID
	if len(args) > 0 {
		if val, ok := args[0].(int); ok {
			workerID = val
		}
	}

	businessID := defaultBusinessID
	if len(args) > 1 {
		if val, ok := args[1].(int); ok {
			businessID = val
		}
	}

	if numShifts == nil {
		defaultNum := 20 // Increase the default number to 20
		numShifts = &defaultNum
	}

	fmt.Printf("WorkerID: %d, BusinessID: %d, NumShifts: %d\n", workerID, businessID, numShifts)
	for i := 0; i < *numShifts; i++ {
		shift := models.Shift{
			StartTime: gofakeit.DateRange(time.Now().AddDate(0, 0, -30), time.Now().AddDate(0, 0, 30)),
			EndTime:   gofakeit.Date(),
			PaymentRate: models.PaymentRate{
				PayAmount: gofakeit.Float64Range(10.0, 50.0),
				PayType:   models.Hourly,
			},
			BusinessID: uint(businessID),
		}

		if err := db.Create(&shift).Error; err != nil {
			return err
		}

		if err := db.Raw(`
			INSERT INTO worker_shifts (worker_id, shift_id)
			VALUES (?, ?)
		`, workerID, shift.ID).Scan(&shift).Error; err != nil {
			return err
		}

		// // Associate the created shift with the worker
		// if err := db.Model(&shift).
		// 	Association("Workers").
		// 	Append(&models.Worker{Model: gorm.Model{ID: uint(workerID)}}); err != nil {
		// 	return err
		// }
	}
	return nil
}

func PopulateReviews(db *gorm.DB, numData *int, args ...interface{}) error {
	var workers []models.Worker
	var businesses []models.Business
	var privateParties []models.PrivateParty

	// Retrieve all workers, businesses, and private parties for the review population.
	if err := db.Find(&workers).Error; err != nil {
		return err
	}
	if err := db.Find(&businesses).Error; err != nil {
		return err
	}
	if err := db.Find(&privateParties).Error; err != nil {
		return err
	}

	if len(workers) == 0 || (len(businesses) == 0 && len(privateParties) == 0) {
		return fmt.Errorf("not enough workers, businesses, or private parties to create reviews")
	}

	for i := 0; i < *numData; i++ {
		// Randomly pick a worker to review.
		worker := workers[gofakeit.Number(0, len(workers)-1)]

		// Randomly decide if a business or private party is leaving the review.
		isBusinessReviewer := true

		review := models.Review{
			Score:    gofakeit.Float32Range(0, 5), // Generate a random score between 0 and 5.
			Comment:  gofakeit.Sentence(10),       // Generate a random sentence of 10 words.
			WorkerID: 1,                           // Associate the review with the selected worker.
			// Below IDs will be set based on the reviewer type chosen.
			BusinessID:     nil,
			PrivatePartyID: nil,
		}

		// Associate with a business or private party based on the random choice.
		if isBusinessReviewer {
			review.BusinessID = &businesses[gofakeit.Number(0, len(businesses)-1)].ID
		} else {
			review.PrivatePartyID = &privateParties[gofakeit.Number(0, len(privateParties)-1)].ID
		}

		// Create the review in the database.
		if err := db.Create(&review).Error; err != nil {
			return err
		}

		// Update the worker's rating after adding the review.
		if err := repo.UpdateWorkerRating(db, worker.ID); err != nil {
			return err
		}
	}

	return nil
}

func PopulateDatabase() error {
	// make gorm print out the SQL statements
	data.DB = data.DB.Debug()
	fmt.Println("Populating database...")
	// wait a few seconds for the database to start
	time.Sleep(2 * time.Second)

	// measure the time it takes to run all the migrations
	start := time.Now()
	for _, runPopulation := range PopulationsAndAmounts {
		if err := runPopulation.Func(data.DB, &runPopulation.Amount, runPopulation.args); err != nil {
			return err
		}
	}
	fmt.Println("Finished populating database in", time.Since(start))

	return nil
}
