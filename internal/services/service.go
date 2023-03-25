package services

import "ValorantWebProject/internal/database"

type Service interface {
	CheckUser(username, password string) (bool, error)
}

type UserService struct {
	repo database.Repository
}

func (s *UserService) CheckUser(username, password string) (bool, error) {
	exists, err := s.repo.UserExists(username, password)
	if err != nil {
		return false, err
	}
	return exists, nil
}
