package middlewares

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/GeorgeHN/email-backend/app/db"
	"github.com/GeorgeHN/email-backend/app/models"
)

type Middleware func(http.Handler) http.Handler

func SetDatabase() *db.DB {
	config := db.Config{
		URI:      "",
		Database: "email",
	}

	return db.NewDBConn(config)

}

// TEST
func Ema(EndPoint http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var wg sync.WaitGroup

		wg.Add(2)

		go func() {
			defer wg.Done()
			fmt.Printf("A new has requested at %v with the id of %s", time.Now(), r.URL.Query().Get("id"))
		}()

		go func() {
			defer wg.Done()
			EndPoint.ServeHTTP(w, r)
		}()

		wg.Wait()
	}
}

// How Many people check the emails as well as what time
func EmailVisit(next http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		emailID := r.URL.Query().Get("id")

		// Store the records on the database
		reg := &models.Reg{
			TimeVisited: time.Now().Local().Unix(),
		}

		d := SetDatabase()

		err := d.InsertEmailVisit(emailID, reg)
		if err != nil {
			return
		}

		next.ServeHTTP(w, r)
	}

}

// Check how many people clicked on the links
