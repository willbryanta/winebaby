package repository

import (
	"database/sql"
	"winebaby/internal/models"
)

func (r *MainRepository) GetWines() ([]models.Wine, error) {
	query := `SELECT id, name, year, manufacturer, region, alcohol_content, serving_temp, serving_size, 
                     serving_size_unit, serving_size_unit: serving_size_unit_abbreviation, 
                     serving_size_unit_description, serving_size_unit_description_abbreviation, 
                     serving_size_unit_description_plural, price, rating, type, colour 
                     FROM wines`
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wines []models.Wine
	for rows.Next() {
		var wine models.Wine
		err := rows.Scan(
			&wine.ID,
			&wine.Name,
			&wine.Year,
			&wine.Manufacturer,
			&wine.Region,
			&wine.AlcoholContent,
			&wine.Price,
			&wine.Rating,
			&wine.Type,
			&wine.Colour,
		)
		if err != nil {
			return nil, err
		}
		wines = append(wines, wine)
	}
	return wines, nil
}

func (r *MainRepository) GetWineByID(id int) (models.Wine, error) {
	query := `SELECT id, name, year, manufacturer, region, alcohol_content, serving_temp, serving_size, 
                     serving_size_unit, serving_size_unit_abbreviation, 
                     serving_size_unit_description, serving_size_unit_description_abbreviation, 
                     serving_size_unit_description_plural, price, rating, type, colour 
                     FROM wines WHERE id = $1`
	var wine models.Wine
	err := r.DB.QueryRow(query, id).Scan(
		&wine.ID,
		&wine.Name,
		&wine.Year,
		&wine.Manufacturer,
		&wine.Region,
		&wine.AlcoholContent,
		&wine.Rating,
		&wine.Type,
		&wine.Colour,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Wine{}, nil
		}
		return models.Wine{}, err
	}
	return wine, nil
}

func (r *MainRepository) CreateWine(wine models.Wine) (int, error) {
	query := `INSERT INTO wines (name, year, manufacturer, region, alcohol_content, serving_temp, 
                               serving_size, serving_size_unit, serving_size_unit_abbreviation, 
                               serving_size_unit_description, serving_size_unit_description_abbreviation, 
                               serving_size_unit_description_plural, price, rating, type, colour) 
              VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16) 
              RETURNING id`
	var id int
	err := r.DB.QueryRow(query,
		wine.Name,
		wine.Year,
		wine.Manufacturer,
		wine.Region,
		wine.AlcoholContent,
		wine.Price,
		wine.Rating,
		wine.Type,
		wine.Colour,
	).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

func (r *MainRepository) UpdateWine(wine models.Wine) error {
	query := `UPDATE wines SET name = $1, year = $2, manufacturer = $3, region = $4, 
                            alcohol_content = $5, serving_temp = $6, serving_size = $7, 
                            serving_size_unit = $8, serving_size_unit_abbreviation = $9, 
                            serving_size_unit_description = $10, serving_size_unit_description_abbreviation = $11, 
                            serving_size_unit_description_plural = $12, price = $13, rating = $14, 
                            type = $15, colour = $16 
              WHERE id = $17`
	_, err := r.DB.Exec(query,
		wine.Name,
		wine.Year,
		wine.Manufacturer,
		wine.Region,
		wine.AlcoholContent,
		wine.Price,
		wine.Rating,
		wine.Type,
		wine.Colour,
		wine.ID,
	)
	return err
}

func (r *MainRepository) DeleteWine(id int) error {
	query := `DELETE FROM wines WHERE id = $1`
	_, err := r.DB.Exec(query, id)
	return err
}