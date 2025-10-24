package factory

import (
	"qr-code-generator/internal/generator"
)

// factory method interface
type QRGenerator interface {
	Generate(data string, opts map[string]string) (string, error)
}

// QRFactory return appropriate generator
func QRFactory(genType string) QRGenerator {
	switch genType {
	case "standard":
		return &generator.StandardQRGenerator{}
	case "custom":
		return &generator.CustomQRGenerator{}
	case "batch":
		return &generator.BatchQRGenerator{}
	default:
		panic("Unsupported generator type")
	}

}
