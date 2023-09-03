package main

import (
	"log"
	"net/http"

	"github.com/Solomon04/go-docker-tutorial/badges"
	"github.com/Solomon04/go-docker-tutorial/companys"
	"github.com/Solomon04/go-docker-tutorial/courses"
	"github.com/Solomon04/go-docker-tutorial/jobs"
	"github.com/Solomon04/go-docker-tutorial/users"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	http.HandleFunc("/users", users.Users)
	http.HandleFunc("/users/badges", users.UserBadges)
	http.HandleFunc("/users/courses", users.UserCourses)
	http.HandleFunc("/badges", badges.Badges)
	http.HandleFunc("/companys", companys.Companys)
	http.HandleFunc("/company/jobs", companys.CompanysJobs)
	http.HandleFunc("/courses", courses.Courses)
	http.HandleFunc("/courses/badges", courses.CourseBadges)
	http.HandleFunc("/jobs", jobs.Jobs)
	http.HandleFunc("/jobs/badges", jobs.JobsBadges)

	// http.HandleFunc("/send-course-workload", user.Workload)

	// http.HandleFunc("/new-user", jobs.JobsBadges)
	// http.HandleFunc("/new-course", jobs.JobsBadges)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
