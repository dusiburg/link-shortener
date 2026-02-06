package services

import (
	"math/rand"

	"example.com/m/internal/repositories"
	"github.com/sqids/sqids-go"
)

type LinksService struct {
	repo *repositories.LinksRepository
	sqid *sqids.Sqids
}

func NewLinksService(repo *repositories.LinksRepository) *LinksService {
	s, _ := sqids.New()
	return &LinksService{
		repo: repo,
		sqid: s,
	}
}

func (s *LinksService) CreateLink(original string) (string, error) {
	id, _ := s.sqid.Encode([]uint64{rand.Uint64()})
	err := s.repo.CreateLink(id, original)
	return id, err
}

func (s *LinksService) GetLink(id string) (string, error) {
	return s.repo.GetLink(id)
}
