package jobs

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/equipe-66-aba/asm/badges"
)

type Job struct {
	ID              int    `json:"id,omitempty"`
	CompanyID       int    `json:"company-id,omitempty"`
	Title           string `json:"title,omitempty"`
	Job_description string `json:"job-description,omitempty"`
	IsTrial         string `json:"temporary-job,omitempty"`
}

type JobBadge struct {
	Job    Job            `json:"job"`
	Badges []badges.Badge `json:"badges"`
}

func Jobs(w http.ResponseWriter, r *http.Request) {
	b := getJobs()

	fmt.Println("Endpoint Hit: jobsPage")
	json.NewEncoder(w).Encode(b)
}

func JobsBadges(w http.ResponseWriter, r *http.Request) {
	jbs := getJobsBadges()

	fmt.Println("Endpoint Hit: JobsBadges")
	json.NewEncoder(w).Encode(jbs)
}

func getJobs() []*Job {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM jobs")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var jobs []*Job
	for results.Next() {
		var j Job
		err = results.Scan(&j.ID, &j.CompanyID, &j.Title, &j.Job_description, &j.IsTrial)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		jobs = append(jobs, &j)
	}

	return jobs
}

func getJobsBadges() []JobBadge {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query(`SELECT j.title, j.job_description, j.istrial, b.name, b.id as idb FROM badges_jobs bj 
								JOIN badges b ON b.id = bj.badgeID
								JOIN jobs j ON j.id = bj.jobsID`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	a := make(map[Job][]badges.Badge)

	for results.Next() {

		var jb Job
		var badge badges.Badge
		err = results.Scan(&jb.Title, &jb.Job_description, &jb.IsTrial, &badge.Name, &badge.ID)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		a[jb] = append(a[jb], badge)
	}

	var jbs []JobBadge
	for k, v := range a {
		jb := JobBadge{
			Job:    k,
			Badges: v,
		}
		jbs = append(jbs, jb)
	}

	return jbs
}
