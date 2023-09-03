package badges

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Badge struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

func Badges(w http.ResponseWriter, r *http.Request) {
	b := getBadges()

	fmt.Println("Endpoint Hit: BadgesPage")
	json.NewEncoder(w).Encode(b)
}

func getBadges() []*Badge {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM badges")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	var badges []*Badge
	for results.Next() {
		var b Badge
		err = results.Scan(&b.ID, &b.Name)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		badges = append(badges, &b)
	}

	return badges
}

func getBadge(ID int) Badge {
	db, err := sql.Open("mysql", "tester:secret@tcp(db:3306)/test")
	if err != nil {
		log.Print(err.Error())
	}
	defer db.Close()

	result := db.QueryRow("SELECT * FROM badges WHERE id=%d", ID)

	var b Badge
	err = result.Scan(&b.ID, &b.Name)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}

	return b
}
