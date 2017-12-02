package main

import (
	"fmt"
	"github.com/mygotest/gormdemo/models"
	"github.com/mygotest/gormdemo/utils"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

func main() {
	db := utils.GormInit()
	db.LogMode(true)
	db.AutoMigrate(new(models.Language), new(models.Movie), new(models.Artist))

	//langjp := &models.Language{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "jp"}
	//db.Create(langjp)
	initData()
	var movie models.Movie = models.Movie{JoinModelBase: models.JoinModelBase{ID: "34c81d41-4c3a-4ea6-abb0-a8628ea1b8a3"}}
	rows := db.Model(&movie).Association("Language").Count()
	fmt.Println(rows)
	movies := db.Model(&models.Artist{JoinModelBase: models.JoinModelBase{ID: "37dceddc-608d-4afe-be49-2c2cfe9c3800"}}).Association("Movies")

	fmt.Println("rows2: ", movies.Count())
	var movieArr []models.Movie
	movies.Find(&movieArr)
	fmt.Println(len(movieArr))

	//
	var perLang []models.Language = make([]models.Language, 0)
	ldb := db.Preload("Language").Find(&perLang, "ID in (?)", []string{"721388c1-4fe9-4a90-91e2-c1eb9983470d", "d8f93735-b106-40f6-a864-b3a6501daab4"})
	fmt.Println(ldb.Value)

	var artists []models.Artist
	db.Model("Artist").Find(&artists, "id in (?) or id = ?", []string{"02b733a1-940d-4376-8892-2339f783c82b", "132c9a7d-6905-4e00-9d92-579fc462ed95"}, "332a6996-0afb-47c9-afd7-ff739109e5ef")
	fmt.Println(artists)

}

func initData() {
	db := utils.GormInit().Begin()
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("-----------error----------")
			fmt.Println(err)
			db.Rollback()
		}
	}()

	lang1 := models.Language{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "tamil"}
	lang2 := models.Language{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "hindi"}
	lang3 := models.Language{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "english"}
	lang4 := models.Language{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "chinese"}
	languages := []models.Language{lang1, lang2, lang3, lang4}
	vals := []interface{}{}
	//sqlStr := "REPLACE INTO gotest_languages (`id`, `name`, `created_at`, `update_time`, `test_time`) VALUES"
	sqlStr := "REPLACE INTO gotest_languages (`id`, `name`, `created_at`, `update_time`) VALUES"
	for _, lang := range languages {
		val := []interface{}{
			lang.ID,
			lang.Name,
			lang.CreatedAt,
			//nil,
			time.Now(),
		}
		q := strings.Repeat("?,", len(val))
		q = q[0 : len(q)-1]
		sqlStr += fmt.Sprintf("(%s),", q)
		vals = append(vals, val...)
	}
	sqlStr = sqlStr[:len(sqlStr)-1]
	fmt.Println(sqlStr)
	fmt.Println(vals)
	fmt.Println(vals...)
	stmt, err := db.CommonDB().Prepare(sqlStr)
	if err != nil {
		panic(err)
	}
	rel, err := stmt.Exec(vals...)
	if err != nil {
		panic(err)
	}
	rows, error := rel.RowsAffected()
	if error != nil {
		panic(error)
	}
	fmt.Println(rows, error)

	//currTime := time.Now()
	movie1 := models.Movie{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Title: "Nayagan", LanguageID: lang2.ID}
	//db.NewRecord(movie1)
	c := db.Create(&movie1)
	error = c.Error
	movie2 := models.Movie{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Title: "Anbe sivam", LanguageID: lang2.ID}
	c = db.Create(&movie2)
	error = c.Error

	movie3 := models.Movie{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Title: "3 idiots", LanguageID: lang3.ID}
	c = db.Create(&movie3)
	error = c.Error

	movie4 := models.Movie{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Title: "Shamithab", LanguageID: lang3.ID}
	c = db.Create(&movie4)
	error = c.Error

	movie5 := models.Movie{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Title: "Dark Knight", LanguageID: lang1.ID}
	c = db.Create(&movie5)
	error = c.Error

	movie6 := models.Movie{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Title: "310 to Yuma", LanguageID: lang1.ID}
	c = db.Create(&movie6)
	error = c.Error

	// argist
	artist1 := models.Artist{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "Madhavan", Movies: []models.Movie{movie2, movie3}}
	db.Create(&artist1)

	artist2 := models.Artist{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "Kamal Hassan", Movies: []models.Movie{movie1, movie2}}
	db.Create(&artist2)

	artist3 := models.Artist{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "Dhanush", Movies: []models.Movie{movie4}}
	db.Create(&artist3)

	artist4 := models.Artist{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "Aamir Khan", Movies: []models.Movie{movie3}}
	db.Create(&artist4)

	artist5 := models.Artist{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "Amitabh Bachchan", Movies: []models.Movie{movie4}}
	db.Create(&artist5)

	artist6 := models.Artist{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "Christian Bale", Movies: []models.Movie{movie5, movie6}}
	db.Create(&artist6)

	artist7 := models.Artist{JoinModelBase: models.JoinModelBase{ID: uuid.NewV4().String(), CreatedAt: time.Now(), TestTime: time.Now()}, Name: "Russell Crowe", Movies: []models.Movie{movie6}}
	db.Create(&artist7)

	db.Commit()
}
