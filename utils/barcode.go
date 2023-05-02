package utils

import (
	"bytes"
	"encoding/base64"
	"image/png"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
)

func GenerateQRCode(content string, size int) (string, error) {
	// Create the QR code
	qrCode, err := qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		return "", err
	}

	// Scale the QR code to the desired size
	qrCode, err = barcode.Scale(qrCode, size, size)
	if err != nil {
		return "", err
	}

	// Encode the QR code as PNG
	var buf bytes.Buffer
	err = png.Encode(&buf, qrCode)
	if err != nil {
		return "", err
	}

	// Convert the PNG image to base64 string
	b64 := base64.StdEncoding.EncodeToString(buf.Bytes())

	return b64, nil
}

func GenerateQRCodev2(content string, filename string, size int) ([]byte, error) {
	// Create the QR code
	qrCode, err := qr.Encode(content, qr.M, qr.Auto)
	if err != nil {
		return nil, err
	}

	// Scale the QR code to the desired size
	qrCode, err = barcode.Scale(qrCode, size, size)
	if err != nil {
		return nil, err
	}

	// Create the output file path
	filePath := "./public/files/" + filename

	// Create the output file directory if it does not exist
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return nil, err
	}

	// Create the output file
	file, err := os.Create(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Encode the QR code as PNG and write to file
	err = png.Encode(file, qrCode)
	if err != nil {
		return nil, err
	}

	// Read the encoded image as bytes
	imageBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	return imageBytes, nil
}
