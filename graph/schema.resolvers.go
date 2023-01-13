package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.22

import (
	"census_server/graph/model"
	"context"
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// EmploymentRate is the resolver for the employmentRate field.
func (r *queryResolver) EmploymentRate(ctx context.Context) (*model.EmploymentRate, error) {
	dsn := "host=containers-us-west-144.railway.app dbname=railway port=5886 sslmode=disable user=postgres password=FUm1gB8FDVcqyXaReo8K"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	// db.AutoMigrate(&model.EmploymentRateData{})
	table := db.Table("censusemployment")

	var LargeResult []*model.EmploymentRateData
	table.Find(&LargeResult)

	source, lastUpdated, nextRelease := "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/lf24/lms", "13/12/2022", "17/1/2023"
	var FinalMeta = &model.Meta{Source: &source, LastUpdated: &lastUpdated, NextRelease: &nextRelease}
	FinalResponse := model.EmploymentRate{Data: LargeResult, Meta: FinalMeta}

	return &FinalResponse, nil
}

// YearlyEmploymentRate is the resolver for the YearlyEmploymentRate field.
func (r *queryResolver) YearlyEmploymentRate(ctx context.Context) (*model.EmploymentRate, error) {
	dsn := "host=containers-us-west-144.railway.app dbname=railway port=5886 sslmode=disable user=postgres password=FUm1gB8FDVcqyXaReo8K"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	// db.AutoMigrate(&model.EmploymentRateData{})
	table := db.Table("yearlyemployment")

	var LargeResult []*model.EmploymentRateData
	table.Find(&LargeResult)

	source, lastUpdated, nextRelease := "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/lf24/lms", "13/12/2022", "17/1/2023"
	var FinalMeta = &model.Meta{Source: &source, LastUpdated: &lastUpdated, NextRelease: &nextRelease}
	FinalResponse := model.EmploymentRate{Data: LargeResult, Meta: FinalMeta}

	return &FinalResponse, nil
}

// MonthlyEmploymentRate is the resolver for the MonthlyEmploymentRate field.
func (r *queryResolver) MonthlyEmploymentRate(ctx context.Context) (*model.EmploymentRate, error) {
	dsn := "host=containers-us-west-144.railway.app dbname=railway port=5886 sslmode=disable user=postgres password=FUm1gB8FDVcqyXaReo8K"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	// db.AutoMigrate(&model.EmploymentRateData{})
	table := db.Table("monthlyemployment")

	var LargeResult []*model.EmploymentRateData
	table.Find(&LargeResult)

	source, lastUpdated, nextRelease := "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/lf24/lms", "13/12/2022", "17/1/2023"
	var FinalMeta = &model.Meta{Source: &source, LastUpdated: &lastUpdated, NextRelease: &nextRelease}
	FinalResponse := model.EmploymentRate{Data: LargeResult, Meta: FinalMeta}

	return &FinalResponse, nil
}

// QuarterlyEmploymentRate is the resolver for the QuarterlyEmploymentRate field.
func (r *queryResolver) QuarterlyEmploymentRate(ctx context.Context) (*model.EmploymentRate, error) {
	dsn := "host=containers-us-west-144.railway.app dbname=railway port=5886 sslmode=disable user=postgres password=FUm1gB8FDVcqyXaReo8K"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	// db.AutoMigrate(&model.EmploymentRateData{})
	table := db.Table("quarterlyemployment")

	var LargeResult []*model.EmploymentRateData
	table.Find(&LargeResult)

	source, lastUpdated, nextRelease := "https://www.ons.gov.uk/employmentandlabourmarket/peopleinwork/employmentandemployeetypes/timeseries/lf24/lms", "13/12/2022", "17/1/2023"
	var FinalMeta = &model.Meta{Source: &source, LastUpdated: &lastUpdated, NextRelease: &nextRelease}
	FinalResponse := model.EmploymentRate{Data: LargeResult, Meta: FinalMeta}

	return &FinalResponse, nil
}

// EmploymentRateYear is the resolver for the EmploymentRateYear field.
func (r *queryResolver) EmploymentRateYear(ctx context.Context, year string) (*model.EmploymentRateData, error) {
	dsn := "host=containers-us-west-144.railway.app dbname=railway port=5886 sslmode=disable user=postgres password=FUm1gB8FDVcqyXaReo8K"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	db.AutoMigrate(&model.EmploymentRate{})
	table := db.Table("censusemployment")

	var result model.EmploymentRateData

	table.Find(&result).Where("timeperiod = ?", year).First(&result)
	// table.Where("timeperiod = ?", year).First(&result)

	return &result, nil
}

// Population is the resolver for the Population field.
func (r *queryResolver) Population(ctx context.Context) ([]*model.Population, error) {
	dsn := "host=containers-us-west-144.railway.app dbname=railway port=5886 sslmode=disable user=postgres password=FUm1gB8FDVcqyXaReo8K"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	db.AutoMigrate(&model.EmploymentRate{})
	table := db.Table("censuspopulation")

	var LargeResult []*model.Population
	table.Find(&LargeResult)

	return LargeResult, nil
}

// PopulationYear is the resolver for the PopulationYear field.
func (r *queryResolver) PopulationYear(ctx context.Context, year string) (*model.Population, error) {
	dsn := "host=containers-us-west-144.railway.app dbname=railway port=5886 sslmode=disable user=postgres password=FUm1gB8FDVcqyXaReo8K"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	db.AutoMigrate(&model.EmploymentRate{})
	table := db.Table("censusPopulation")

	var result model.Population
	table.First(&result)
	table.Where("timeperiod = ?", year).First(&result)
	// fmt.Println(result.TimePeriod)
	return &result, nil
}

// GenderIdentity is the resolver for the GenderIdentity field.
func (r *queryResolver) GenderIdentity(ctx context.Context) ([]*model.GenderIdentity, error) {
	dsn := "host=containers-us-west-144.railway.app dbname=railway port=5886 sslmode=disable user=postgres password=FUm1gB8FDVcqyXaReo8K"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("connected to database")
	}
	db.AutoMigrate(&model.EmploymentRate{})
	table := db.Table("censusgender")

	var LargeResult []*model.GenderIdentity
	table.Find(&LargeResult)

	return LargeResult, nil
}

// Ping is the resolver for the ping field.
func (r *queryResolver) Ping(ctx context.Context) (string, error) {
	return "pong", nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
type mutationResolver struct{ *Resolver }