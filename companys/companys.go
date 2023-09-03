package companys

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/equipe-66-aba/asm/jobs"
)

type Company struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type CompanyJobs struct {
	CompanyName string     `json:"companyName"`
	Jobs        []jobs.Job `json:"jobs"`
}

type JobBadge struct {
	Jobs []jobs.Job `json:"jobs"`
}

func Companys(w http.ResponseWriter, r *http.Request) {
	b := getCompanys()

	fmt.Println("Endpoint Hit: companysPage")
	json.NewEncoder(w).Encode(b)
}

func CompanysJobs(w http.ResponseWriter, r *http.Request) {
	cJobs := getCompanysJobs()

	fmt.Println("Endpoint Hit: CompanysJobs")
	json.NewEncoder(w).Encode(cJobs)
}

func getCompanys() []*Company {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM companys")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var companys []*Company
	for results.Next() {
		var b Company
		err = results.Scan(&b.ID, &b.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		companys = append(companys, &b)
	}

	return companys
}

func getCompanysJobs() []CompanyJobs {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query(`SELECT c.name, j.title, j.job_description, j.istrial FROM jobs j 
								JOIN companys c ON c.id = j.companyID`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	a := make(map[string][]jobs.Job)

	for results.Next() {

		var cj CompanyJobs
		var job jobs.Job
		err = results.Scan(&cj.CompanyName, &job.Title, &job.Job_description, &job.IsTrial)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		a[cj.CompanyName] = append(a[cj.CompanyName], job)
	}

	var cjs []CompanyJobs
	for k, v := range a {
		cj := CompanyJobs{
			CompanyName: k,
			Jobs:        v,
		}
		cjs = append(cjs, cj)
	}

	return cjs
}
