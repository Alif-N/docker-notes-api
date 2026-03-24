package service

import (
	"errors"
	"notes-api/model"
	"notes-api/repository"
)

func CreateNote(note *model.Note) error {
	if note.Title == "" {
		return errors.New("title is required")
	}
	if note.Content == "" {
		return errors.New("content is required")
	}
	return repository.CreateNote(note)
}

func GetNotes() ([]model.Note, error) {
	return repository.GetNotes()
}

func GetNoteByID(id string) (*model.Note, error) {
	return repository.GetNoteByID(id)
}

func UpdateNote(id string, note *model.Note) error {
	if note.Title == "" {
		return errors.New("title is required")
	}
	if note.Content == "" {
		return errors.New("content is required")
	}
	return repository.UpdateNote(id, note)
}

func DeleteNote(id string) (int64, error) {
	return repository.DeleteNote(id)
}
