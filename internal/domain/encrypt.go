package domain

type EncryptedData struct {
	Input  string `json:"input"`
	MD5    string `json:"md5"`
	SHA256 string `json:"sha256"`
}

type EncryptRequest struct {
	Input string `json:"input"`
}

type EncryptRepository interface {
	Save(data EncryptedData) error
	Get(hash string) (*EncryptedData, error)
}
