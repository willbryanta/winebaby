package repository

import "winebaby/models"

var reviews = []models.Review{
	{ID: 1, WineID:1, Author:"Grant Burge", Description: "Zesty and fruity with notes of licorice", Rating: 7}
}

// Get all reviews
func GetReviews() []models.Review{
	return reviews
}

// Get a review by ID
func GetReviewById(id int) *modles.Review{
	for _,r := range reviews{
		if r.ID == id{
			return &r
		}
	}
	return nil
}

func CreateReview(r models.Review){
	reviews = append(reviews, r)
}

func UpdateReview(id int, updated models.Review) bool {
	for i, r := range reviews{
		if r.ID == id {
			reviews[i] = updated
			return true
		}
	}
	return false
}

func DeleteReview(id int) bool {
	for i, r:= range reviews {
		if r.ID ==id {
			reviews = append(redviews[:i], reviews[i+1:]...)
			return true
		}
	}
	return false
}