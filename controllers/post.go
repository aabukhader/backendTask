package controllers

import (
	"database/sql"
	"encoding/json"
	"log"
	"math"
	"net/http"
	"strconv"

	"github.com/aabukhader/backEnd/db"
	"github.com/aabukhader/backEnd/helper"
	"github.com/aabukhader/backEnd/models"
)

var connic *sql.DB = db.ConnectDb()

func insertPost(post *models.PostItem) error {
	stmt, _ := connic.Prepare("INSERT INTO post (name, type, url, global_id, description, published_at, publisher ) VALUES(?,?,?,?,?,?,?)")
	_, err := stmt.Exec(post.Name, post.Type, post.URL, post.GlobalID, post.Description, post.PublishedAt, post.Publisher)
	defer stmt.Close()
	return err
}

// Post function
func Post(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var post *models.PostItem
	json.NewDecoder(r.Body).Decode(&post)
	err := insertPost(post)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		var res models.StatusRes
		res.Status = 500
		res.Msg = "Something wont wrong"
		json.NewEncoder(w).Encode(res)
	} else {
		w.WriteHeader(http.StatusOK)
		var res models.StatusRes
		res.Status = 200
		res.Msg = "New Post has been added Successfully"
		json.NewEncoder(w).Encode(res)
	}
}

func getAllPosts() int {
	var total int
	err := connic.QueryRow("SELECT COUNT(*) FROM post").Scan(&total)
	if err != nil {
		log.Fatal(err)
	}
	return total
}

func paginatePosts(begin, limit int) ([]models.PostItem, error) {
	var items []models.PostItem
	res, err := connic.Query("SELECT * FROM post LIMIT ? OFFSET ?", limit, begin)

	for res.Next() {
		var item models.PostItem
		res.Scan(
			&item.ID,
			&item.Name,
			&item.Type,
			&item.URL,
			&item.GlobalID,
			&item.Description,
			&item.Publisher,
			&item.PublishedAt)
		items = append(items, item)
	}
	return items, err
}

// GetPosts function
func GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	total := getAllPosts()
	limit := 5
	page, begin := helper.Pagination(r, limit)
	pages := (total / limit)
	totalPages := math.Ceil(float64(total) / float64(limit))
	var nextPage int
	if page+1 <= int(totalPages) {
		nextPage = page + 1
	}
	if (total % limit) != 0 {
		pages++
		nextPage = page
	}

	result, err := paginatePosts(begin, limit)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		var res models.StatusRes
		res.Status = 500
		res.Msg = "Something wont wrong"
		json.NewEncoder(w).Encode(res)
	} else {
		w.WriteHeader(http.StatusOK)
		var res models.PageStatusResSuccss
		res.Status = 200
		res.Msg = "Success"
		var meta models.Pagination
		meta.ItemPerPage = limit
		if page+1 > int(totalPages) {
			meta.NextPageURL = ""
		} else {
			meta.NextPageURL = "/api/post/getAll?page=" + strconv.Itoa(nextPage)
		}
		if page-1 != 0 {
			prePage := page - 1
			meta.PrevPageURL = "/api/post/getAll?page=" + strconv.Itoa(prePage)
		} else {
			meta.PrevPageURL = ""
		}
		meta.TotalPageNum = totalPages
		meta.TotalCount = total
		meta.Page = page
		res.Data.Meta = meta
		res.Data.Items = result
		json.NewEncoder(w).Encode(res)
	}
}
