package audience

import (
	"encoding/csv"
	"net/http"

	"github.com/GeorgeHN/email-backend/app/models"
)

func ProcessAudience(url string) ([]*models.Audience, error) {

	csv, err := downloadCSV(url)
	if err != nil {
		return nil, err
	}

	au := make([]*models.Audience, 0, len(csv)-1) // Exclude the header row

	for i, record := range csv {
		if i == 0 { // Skip the header row
			continue
		}

		if len(record) < 3 { // Ensure the record has at least three fields
			continue
		}

		audience := models.Audience{
			First: record[0],
			Last:  record[1],
			Email: record[2],
		}
		au = append(au, &audience)
	}

	return au, nil
}

func downloadCSV(url string) ([][]string, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	reader := csv.NewReader(response.Body)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	return records, nil
}
