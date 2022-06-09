package service

import "github.com/ShatteredRealms/Characters/pkg/repository"

type CharacterService interface {
}

type characterService struct {
	characterRepository repository.CharacterRepository
}

func NewCharacterService(r repository.CharacterRepository) CharacterService {
	return characterService{
		characterRepository: r,
	}
}
