package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Couse struct {
	CouseID string  `json:"courseid"`
	Title   string  `json:"title"`
	Price   int     `json:"price"`
	Author  *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake DB
var courses []Couse

// middlewares
func (c *Couse) IsEmpty() bool {
	return c.Title == "" && c.Price == 0 && c.Author == nil
}
func main() {

}

//	app.get("/", (req, res) => {
//	    res.send("Welcome to API home page");
//	});
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to API home page"))
}

//	app.get("/courses", (req, res) => {
//	    res.json(courses);
//	});
func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

// app.get("/course/:id", (req, res) => {
//     res.set("Content-Type", "application/json");

//     const { id } = req.params;

//     for (const course of courses) {
//         if (course.courseid === id) {
//             return res.json(course);
//         }
//     }

//     res.status(404).json({ error: "No course found with given id" });
// });

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// grab id from request
	params := mux.Vars(r)
	fmt.Println(params)
	fmt.Print("ayush")
	// loop through courses, find matching id

	for _, courses := range courses {
		if courses.CouseID == params["id"] {
			json.NewEncoder(w).Encode(courses)
			return
		}
	}
	fmt.Errorf("No course found with given id")
	fmt.Println("Get one course")
}
