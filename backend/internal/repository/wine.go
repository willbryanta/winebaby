package repository

import (
	"database/sql"
	"winebaby/internal/models"
)

func (r *Repository) GetWines() ([]models.Wine, error) {
	query := `SELECT id, name, region, varietal FROM wines`
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var wines []models.Wine
	for rows.Next() {
		var wine models.Wine
		if err := rows.Scan(&wine.ID, &wine.Name, &wine.Region, &wine.Varietal); err != nil {
			return nil, err
		}
		wines = append(wines, wine)
	}
	return wines, nil
}
func (r *Repository) GetWineByID(id int) (models.Wine, error) {
	query := `SELECT id, name, region, varietal FROM wines WHERE id = $1`
	var wine models.Wine
	err := r.db.QueryRow(query, id).Scan(&wine.ID, &wine.Name, &wine.Region, &wine.Varietal)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Wine{}, nil
		}
		return models.Wine{}, err
	}
	return wine, nil
}
func (r *Repository) CreateWine(wine models.Wine) (int, error) {
	query := `INSERT INTO wines (name, region, varietal) VALUES ($1, $2, $3) RETURNING id`
	err := r.db.QueryRow(query, wine.Name, wine.Region, wine.Varietal).Scan(&wine.ID)
	if err != nil {
		return 0, err
	}
	return wine.ID, nil
}
func (r *Repository) UpdateWine(wine models.Wine) error {
	query := `UPDATE wines SET name = $1, region = $2, varietal = $3 WHERE id = $4`
	_, err := r.db.Exec(query, wine.Name, wine.Region, wine.Varietal, wine.ID)
	return err
}
func (r *Repository) DeleteWine(id int) error {
	query := `DELETE FROM wines WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}