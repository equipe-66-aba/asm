package main

import (
	"log"
	"net/http"

	"github.com/equipe-66-aba/asm/badges"
	"github.com/equipe-66-aba/asm/companys"
	"github.com/equipe-66-aba/asm/courses"
	"github.com/equipe-66-aba/asm/jobs"
	"github.com/equipe-66-aba/asm/users"
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
	http.HandleFunc("/courses/tracks", courses.CourseTracks)
	http.HandleFunc("/jobs", jobs.Jobs)
	http.HandleFunc("/jobs/badges", jobs.JobsBadges)

	// http.HandleFunc("/send-course-workload", user.Workload)
	// http.HandleFunc("/new-user", user.New)
	// http.HandleFunc("/new-course", course.New)
	// http.HandleFunc("/new-company", company.New)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
