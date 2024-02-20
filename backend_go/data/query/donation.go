package query

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"zbackend/data/models"
)

func readDonationFile() *os.File {
	file, err := os.Open("data/100-last-donations.json")
	if err != nil {
		log.Fatal(err)
	}
	return file
}

func readFile(reader io.Reader) ([]byte, error) {
	bytes, err := io.ReadAll(reader)
	return bytes, err
}

func GetDonations() ([]models.TransactionWithDonationObject, error) {
	file := readDonationFile()
	defer file.Close()

	bytes, err := readFile(file)
	if err != nil {
		return nil, err
	}

	return parseDonations(bytes)
}

func GetDonationsPaginated(limit int, offset int) []models.TransactionWithDonationObject {
	donations, err := GetDonations()
	if err != nil {
		log.Fatal(err)
	}

	return sliceDonationsPaginated(donations, limit, offset)
}

func parseDonations(bytes []byte) ([]models.TransactionWithDonationObject, error) {
	var donations []models.TransactionWithDonationObject
	if err := json.Unmarshal(bytes, &donations); err != nil {
		return nil, err
	}

	return donations, nil
}

func sliceDonationsPaginated(donations []models.TransactionWithDonationObject, limit int, offset int) []models.TransactionWithDonationObject {
	startIndex := max(offset, 0)
	endIndex := min(offset+limit, len(donations))

	// slice donations
	return donations[startIndex:endIndex]
}
