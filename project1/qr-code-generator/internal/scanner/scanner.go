package scanner

import (
	"fmt"
	"os"
	"qr-code-generator/internal/db"
	"qr-code-generator/internal/models"

	"github.com/tuotoo/qrcode"
)

func Scan(filePath string) (string, error) {
	fi, err := os.Open(filePath)
	if err != nil {
		return "", fmt.Errorf("Failed to open file: %w", err)
	}

	defer fi.Close()

	qr, err := qrcode.Decode(fi)
	if err != nil {
		return "", fmt.Errorf("Failed to decode qrcode : %w", err)
	}

	record := models.QRRecord{Data: qr.Content, Type: "Scanned", FilePath: filePath}
	if db.DB != nil {
		db.DB.Create(&record)
	}

	return qr.Content, nil
}
