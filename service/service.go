package service

import (
	"devnotes/db"
	"devnotes/model"
	"devnotes/storage"

	"github.com/google/uuid"
)

type Service struct {
	store *storage.MemoryStore
}

func NewService(store *storage.MemoryStore) *Service {
	return &Service{store: store}
}

func (s *Service) Register(username, password string) (model.User, error) {
	id := uuid.New().String()
	user := model.User{
		ID:       id,
		Username: username,
		Password: password,
	}

	s.store.Mu.Lock()
	defer s.store.Mu.Unlock()
	s.store.Users[id] = user

	return user, nil
}

func (s *Service) CreateNote(userID, title, content string) (model.Note, error) {
	note := model.Note{
		ID:      uuid.New().String(),
		OwnerID: userID,
		Title:   title,
		Content: content,
	}

	s.store.Mu.Lock()
	defer s.store.Mu.Unlock()
	s.store.Notes[note.ID] = note

	return note, nil
}

func (s *Service) GetNotesByUserID(userID string) []model.Note {
	var userNotes []model.Note
	for _, note := range s.store.Notes {
		if note.OwnerID == userID {
			userNotes = append(userNotes, note)
		}
	}
	return userNotes
}

func (s *Service) UpdateNote(noteID, title, content string) (model.Note, error) {
	var note model.Note
	if err := db.DB.First(&note, "id = ?", noteID).Error; err != nil {
		return model.Note{}, err
	}

	note.Title = title
	note.Content = content

	if err := db.DB.Save(&note).Error; err != nil {
		return model.Note{}, err
	}

	return note, nil
}

func (s *Service) DeleteNote(noteID string) error {
	return db.DB.Delete(&model.Note{}, "id = ?", noteID).Error
}
