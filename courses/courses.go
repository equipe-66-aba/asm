package courses

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Course struct {
	ID        int    `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	CompanyID int    `json:"companyId,omitempty"`
	Workload  int    `json:"workload,omitempty"`
}

type CourseBadge struct {
	CourseName string   `json:"courseName"`
	BadgesName []string `json:"badges"`
}

type CourseTrack struct {
	CourseName string `json:"courseName"`
	MainTrack  string `json:"maintrack"`
	SubTrack   string `json:"subtrack"`
}

func Courses(w http.ResponseWriter, r *http.Request) {
	courses := getCourses()

	fmt.Println("Endpoint Hit: courses")
	json.NewEncoder(w).Encode(courses)
}

func CourseBadges(w http.ResponseWriter, r *http.Request) {
	cb := getCoursesBadges()

	fmt.Println("Endpoint Hit: coursesBadges")
	json.NewEncoder(w).Encode(cb)
}

func CourseTracks(w http.ResponseWriter, r *http.Request) {
	ct := getCoursesTrack()

	fmt.Println("Endpoint Hit: CourseTracks")
	json.NewEncoder(w).Encode(ct)
}

func getCourses() []*Course {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM coursers")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var courses []*Course
	for results.Next() {
		var c Course
		err = results.Scan(&c.ID, &c.Name, &c.CompanyID, &c.Workload)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		courses = append(courses, &c)
	}

	return courses
}

func getCoursesBadges() []CourseBadge {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query(`SELECT c.name, b.name FROM badges_coursers cb 
								JOIN badges b ON b.id = cb.badgeID
								JOIN coursers c ON c.id = cb.courserID`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	a := make(map[string][]string)

	for results.Next() {

		var c CourseBadge
		var b string
		err = results.Scan(&c.CourseName, &b)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		a[c.CourseName] = append(a[c.CourseName], b)
	}

	var cbs []CourseBadge
	for k, v := range a {
		cb := CourseBadge{
			CourseName: k,
			BadgesName: v,
		}
		cbs = append(cbs, cb)
	}

	return cbs
}

func getCoursesTrack() []CourseTrack {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query(`SELECT c.name, t.name, st.name FROM courses_sub_tracks cst 
								JOIN sub_tracks st ON st.id = cst.subTrackID AND st.trackID = cst.trackID
								JOIN tracks t ON t.id = st.trackID
								JOIN coursers c ON c.id = cst.courseID`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var cts []CourseTrack
	for results.Next() {

		var c CourseTrack
		err = results.Scan(&c.CourseName, &c.MainTrack, &c.SubTrack)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		cts = append(cts, c)
	}

	return cts
}
