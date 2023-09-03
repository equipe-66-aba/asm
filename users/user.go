package users

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type UserBadge struct {
	UserName   string   `json:"userName"`
	BadgesName []string `json:"badges"`
}

type UserCourse struct {
	UserName string   `json:"userName"`
	Course   []Course `json:"courses"`
}

type Course struct {
	CourseName        string `json:"name"`
	TotalWorkload     int    `json:"totalWorkload"`
	WorkloadCompleted int    `json:"workloadCompleted"`
	PercentCompleted  int    `json:"percentCompleted"`
}

func Users(w http.ResponseWriter, r *http.Request) {
	users := getUsers()

	fmt.Println("Endpoint Hit: Users")
	json.NewEncoder(w).Encode(users)
}

func UserBadges(w http.ResponseWriter, r *http.Request) {
	users := getUsersBadges()

	fmt.Println("Endpoint Hit: UsersBadges")
	json.NewEncoder(w).Encode(users)
}

func UserCourses(w http.ResponseWriter, r *http.Request) {
	users := getUserCourses()

	fmt.Println("Endpoint Hit: UsersBadges")
	json.NewEncoder(w).Encode(users)
}

func getUsers() []*User {
	// Open up our database connection.
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	// Execute the query
	results, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var users []*User
	for results.Next() {
		var u User
		// for each row, scan the result into our tag composite object
		err = results.Scan(&u.ID, &u.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		users = append(users, &u)
	}

	return users
}

func getUsersBadges() []UserBadge {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query(`SELECT u.name, b.name FROM badges_users bu 
								JOIN badges b ON b.id = bu.badgeID
								JOIN users u ON u.id = bu.userID`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	a := make(map[string][]string)

	for results.Next() {

		var u UserBadge
		var b string
		err = results.Scan(&u.UserName, &b)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		a[u.UserName] = append(a[u.UserName], b)
	}

	var us []UserBadge
	for k, v := range a {
		cb := UserBadge{
			UserName:   k,
			BadgesName: v,
		}
		us = append(us, cb)
	}

	return us
}

func getUserCourses() []UserCourse {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")

	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query(`SELECT u.name, c.name, c.workload, uc.workloadCompleted FROM users_coursers uc 
								JOIN coursers c ON c.id = uc.courserID
								JOIN users u ON u.id = uc.userID`)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	a := make(map[string][]Course)

	for results.Next() {

		var uc UserCourse
		var c Course
		err = results.Scan(&uc.UserName, &c.CourseName, &c.TotalWorkload, &c.WorkloadCompleted)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		c.PercentCompleted = int((float64(c.WorkloadCompleted) / float64(c.TotalWorkload)) * 100)
		a[uc.UserName] = append(a[uc.UserName], c)
	}

	var ucs []UserCourse
	for k, v := range a {
		uc := UserCourse{
			UserName: k,
			Course:   v,
		}
		ucs = append(ucs, uc)
	}

	return ucs
}
