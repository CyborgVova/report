package server

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/cyborgvova/report/database"
)

type Answer struct {
	ErrorMsg string `json:"error_msg"`
}

type Answer1 struct {
	ErrorMsg   string `json:"error_msg"`
	ReportInfo string `json:"report_info"`
}

type Answer2 struct {
	ErrorMsg             string `json:"error_msg"`
	MaxObservationPeriod string `json:"max_observation_period"`
}

func getMaxObservationPeriod(w http.ResponseWriter, r *http.Request) {
	answer := &Answer2{}
	id, err := strconv.Atoi(r.URL.Query().Get("model_id"))
	if err != nil {
		answer.ErrorMsg = fmt.Sprintf("%v", err)
	}
	db := database.DBConnect()
	count := 0
	db.Raw("SELECT COUNT(model_id) FROM reports WHERE model_id = ?", id).Scan(&count)
	var days string
	if count == 1 {
		db.Raw("SELECT (current_date - creation_time) AS days "+
			"FROM reports WHERE reports.model_id = ? ", id).Scan(&days)
	} else {
		db.Raw("SELECT COALESCE(creation_time - LAG(creation_time) "+
			"OVER(ORDER BY creation_time), 0) AS days "+
			"FROM reports WHERE model_id = ? "+
			"ORDER BY days DESC LIMIT 1", id).Scan(&days)
	}

	answer.MaxObservationPeriod = days
	b, err := json.Marshal(answer)
	if err != nil {
		answer.ErrorMsg = fmt.Sprintf("%v", err)
	}
	fmt.Fprint(w, string(b))
}

func getReport(w http.ResponseWriter, r *http.Request) {
	answer := &Answer1{}
	id, err := strconv.Atoi(r.URL.Query().Get("report_id"))
	if err != nil {
		answer.ErrorMsg = fmt.Sprintf("%v", err)
	}
	db := database.DBConnect()
	report := &database.Report{}
	if err := db.Select("report_info").
		Where("report_id = ?", id).
		Find(report).Error; err != nil {
		answer.ErrorMsg = fmt.Sprintf("%v", err)
	}
	answer.ReportInfo = report.ReportInfo
	b, err := json.Marshal(answer)
	if err != nil {
		log.Println("error serialization:", err)
	}
	fmt.Fprint(w, string(b))
}

func setReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		return
	}
	answer := &Answer{}
	report := &database.Report{}
	b, err := io.ReadAll(r.Body)
	if err != nil {
		answer.ErrorMsg = fmt.Sprintf("%v", err)
		fmt.Fprint(w, answer)
		return
	}
	if err := json.Unmarshal(b, report); err != nil {
		answer.ErrorMsg = fmt.Sprintf("%v", err)
		fmt.Fprint(w, answer)
		return
	}
	report.CreationTime = time.Now()
	db := database.DBConnect()
	db.Create(report)

	fmt.Fprint(w, answer)
}

func NewServer() {
	http.HandleFunc("/api/v1/get_report", getReport)
	http.HandleFunc("/api/v1/set_report", setReport)
	http.HandleFunc("/api/v1/get_observation_time", getMaxObservationPeriod)

	fmt.Println("Starting server on :8080 ...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Println("error server start:", err)
	}
}
