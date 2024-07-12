package usecase

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"test-international-trade/internal/domain"
)

type EncryptUseCase struct {
	repo domain.EncryptRepository
}

func NewEncryptUseCase(repo domain.EncryptRepository) *EncryptUseCase {
	return &EncryptUseCase{repo: repo}
}

func (uc *EncryptUseCase) Encrypt(input string) (*domain.EncryptedData, error) {

	cache, _ := uc.repo.Get(input)
	if cache != nil {
		return cache, nil
	}

	md5Hash := md5.Sum([]byte(input))
	sha256Hash := sha256.Sum256([]byte(input))

	data := &domain.EncryptedData{
		Input:  input,
		MD5:    hex.EncodeToString(md5Hash[:]),
		SHA256: hex.EncodeToString(sha256Hash[:]),
	}

	if err := uc.repo.Save(*data); err != nil {
		return nil, err
	}

	return data, nil
}
