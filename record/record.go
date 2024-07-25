package record

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

var header = []string{"Date", "Mode", "Speed", "Accuracy"}

type Record struct {
	Date     time.Time
	Mode     string
	Speed    int
	Accuracy float64
}

func getPaths() string {

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		return ""
	}

	configDir := filepath.Join(homeDir, ".config/typeTest-go")
	recordPath := filepath.Join(configDir, "records.json")

	return recordPath
}

func Save(mode string, speed int, accuracy float64) {

	record := &Record{
		Date:     time.Now(),
		Mode:     mode,
		Speed:    speed,
		Accuracy: accuracy,
	}
	recordPath := getPaths()

	file, err := os.OpenFile(recordPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header if the file is new
	if stat, err := file.Stat(); err == nil && stat.Size() == 0 {
		if err := writer.Write(header); err != nil {
			fmt.Println("Error writing header:", err)
			return
		}
	}

	// Prepare record data
	recordData := []string{
		record.Date.Format(time.RFC3339), // Format date to a string
		record.Mode,
		fmt.Sprintf("%d", record.Speed),      // Format speed to a string
		fmt.Sprintf("%.2f", record.Accuracy), // Format accuracy to a string
	}

	// Write record to CSV
	if err := writer.Write(recordData); err != nil {
		fmt.Println("Error writing record:", err)
		return
	}

}
func ReadCSV() ([]Record, error) {
	// Open the CSV file
	recordFile := getPaths()
	file, err := os.Open(recordFile)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the header
	_, err = reader.Read()
	if err != nil {
		return nil, fmt.Errorf("error reading header: %v", err)
	}

	// Read all records
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("error reading records: %v", err)
	}

	// Parse records into a slice of Record structs
	var result []Record
	for _, record := range records {
		date, err := time.Parse(time.RFC3339, record[0])
		if err != nil {
			return nil, fmt.Errorf("error parsing date: %v", err)
		}
		speed, err := strconv.Atoi(record[2])
		if err != nil {
			return nil, fmt.Errorf("error parsing speed: %v", err)
		}
		accuracy, err := strconv.ParseFloat(record[3], 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing accuracy: %v", err)
		}
		result = append(result, Record{
			Date:     date,
			Mode:     record[1],
			Speed:    int(speed),
			Accuracy: accuracy,
		})
	}

	return result, nil
}
