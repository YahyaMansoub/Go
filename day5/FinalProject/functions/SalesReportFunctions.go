package functions

import (
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var reportMutex sync.RWMutex

type SalesReport struct {
	Date            time.Time `json:"date"`
	TotalRevenue    float64   `json:"total_revenue"`
	TotalOrders     int       `json:"total_orders"`
	TotalBooksSold  int       `json:"total_books_sold"`
	TopSellingBooks []string  `json:"top_selling_books"`
}

func GetSalesReportsInRange(startDate, endDate time.Time) ([]SalesReport, error) {
	reportMutex.RLock()
	defer reportMutex.RUnlock()

	var reports []SalesReport

	reportFiles, err := filepath.Glob("output-reports/report_*.json")
	if err != nil {
		return nil, errors.New("failed to read report directory")
	}

	for _, file := range reportFiles {
		content, err := os.ReadFile(file)
		if err != nil {
			return nil, err
		}

		var report SalesReport
		if err := json.Unmarshal(content, &report); err != nil {
			return nil, err
		}

		if (report.Date.After(startDate) || report.Date.Equal(startDate)) &&
			(report.Date.Before(endDate) || report.Date.Equal(endDate)) {
			reports = append(reports, report)
		}
	}

	return reports, nil
}
func GenerateDailySalesReport() error {
	reportMutex.Lock()
	defer reportMutex.Unlock()

	report := SalesReport{
		Date:            time.Now(),
		TotalRevenue:    10000.0,
		TotalOrders:     120,
		TotalBooksSold:  300,
		TopSellingBooks: []string{"Book A", "Book B", "Book C"},
	}

	outputDir := "output-reports"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	reportFileName := filepath.Join(outputDir, "report_"+time.Now().Format("02012006015004")+".json")
	fileContent, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(reportFileName, fileContent, 0644); err != nil {
		return err
	}

	return nil
}
