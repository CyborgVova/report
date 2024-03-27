package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/cyborgvova/report/database"
)

type Answer1 struct {
	ErrorMsg   string `json:"error_msg"`
	ReportInfo string `json:"report_info"`
}

type Answer2 struct {
	ErrorMsg             string `json:"error_msg"`
	MaxObservationPeriod string `json:"max_observation_period"`
}

func getReport(w http.ResponseWriter, r *http.Request) {
	answer := &Answer1{}
	id, err := strconv.Atoi(r.URL.Query().Get("report_id"))
	if err != nil {
		answer.ErrorMsg = fmt.Sprintf("%v", err)
	}
	db := database.DBConnect()
	report := &database.Report{}
	if err := db.Where("report_id = ?", id).Find(report).Error; err != nil {
		answer.ErrorMsg = fmt.Sprintf("%v", err)
	}
	answer.ReportInfo = report.ReportInfo

	b, err := json.Marshal(answer)
	if err != nil {
		log.Println("error serialization:", err)
	}
	fmt.Fprint(w, string(b))
}

func NewServer() {
	http.HandleFunc("/api/v1/get_report", getReport)

	fmt.Println("Starting server on :8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("error server start:", err)
	}
}
