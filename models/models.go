package models

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Metadata struct {
	Hash string `json:"Hash"`
	Type string `json:"Type"`
}

// IsDB menyimpan metadata C2 fingerprint yang dimuat dari metadata.json
var IsDB []Metadata

// LoadMetadata membaca metadata dari file JSON. Path default: rat-fingerprint/metadata.json.
// Jika path kosong, menggunakan "rat-fingerprint/metadata.json".
func LoadMetadata(path string) error {
	if path == "" {
		path = "rat-fingerprint/metadata.json"
	}
	data, err := os.ReadFile(path)
	if err != nil {
		// coba path relatif terhadap executable
		exe, _ := os.Executable()
		dir := filepath.Dir(exe)
		path = filepath.Join(dir, "rat-fingerprint", "metadata.json")
		data, err = os.ReadFile(path)
		if err != nil {
			return err
		}
	}
	return json.Unmarshal(data, &IsDB)
}
