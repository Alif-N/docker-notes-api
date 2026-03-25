package repository

import (
	"fmt"
	"notes-api/db"
	"notes-api/model"
)

func CreateNote(note *model.Note) error {
	query := `INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id, created_at, updated_at`
	err := db.DB.QueryRow(query, note.Title, note.Content).Scan(
		&note.ID, 
		&note.CreatedAt, 
		&note.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to create note: %w", err)
	}
	
	return nil
}

func GetNotes(page, limit int, search string) ([]model.Note, int, error) {
	offset := (page - 1) * limit
	
	query := `
		SELECT id, title, content, created_at, updated_at
		FROM notes
		WHERE title ILIKE $1 OR content ILIKE $1
		ORDER BY created_at DESC
		LIMIT $2 OFFSET $3
	`
	rows, err := db.DB.Query(query, "%"+search+"%", limit, offset)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find notes: %w", err)
	}
	defer rows.Close()

	var notes []model.Note
	for rows.Next() {
		var note model.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content, &note.CreatedAt, &note.UpdatedAt); err != nil {
			return nil, 0, fmt.Errorf("failed to scan note: %w", err)
		}
		notes = append(notes, note)
	}

	var total int
	countQuery := `
		SELECT COUNT(*)
		FROM notes
		WHERE title ILIKE $1 OR content ILIKE $1
	`
	err = db.DB.QueryRow(countQuery, "%"+search+"%").Scan(&total)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count notes: %w", err)
	}

	return notes, total, nil
}

func GetNoteByID(id string) (*model.Note, error) {
	var note model.Note
	query := `SELECT id, title, content, created_at, updated_at FROM notes WHERE id = $1`
	err := db.DB.QueryRow(query, id).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch note: %w", err)
	}

	return &note, nil
}

func UpdateNote(id string, note *model.Note) error {
	query := `UPDATE notes SET title = $1, content = $2, updated_at = NOW() 
			  WHERE id = $3 RETURNING id, title, content, created_at, updated_at`
	err := db.DB.QueryRow(query, note.Title, note.Content, id).Scan(
		&note.ID,
		&note.Title,
		&note.Content,
		&note.CreatedAt,
		&note.UpdatedAt,
	)
	if err != nil {
		return fmt.Errorf("failed to update note: %w", err)
	}
	return nil
}

func DeleteNote(id string) (int64, error) {
	result, err := db.DB.Exec(`DELETE FROM notes WHERE id = $1`, id)
	if err != nil {
		return 0, fmt.Errorf("failed to delete note: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("failed to get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return 0, fmt.Errorf("note with ID %s not found", id)
	}

	return rowsAffected, nil
}
