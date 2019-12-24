package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"log"
	"strconv"
)

//Student This is student type
type Student struct {
	Name      string
	Branch    string
	StudentID int
	Score     float64
}

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	if _, err := conn.Do("AUTH", "infracloud"); err != nil {
		conn.Close()
	}
	_, err = conn.Do("HMSET", "student:1", "Name", "Mahendra", "Branch", "IT", "StudentID", 2010, "Score", 8.4344)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Student added!")

	reply, err := redis.StringMap(conn.Do("HGETALL", "student:1"))
	if err != nil {
		log.Fatal(err)
	}
	student, err := populateStudent(reply)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", student)
}

func populateStudent(reply map[string]string) (*Student, error) {
	var err error
	student := new(Student)
	student.Name = reply["Name"]
	student.Branch = reply["Branch"]
	student.StudentID, err = strconv.Atoi(reply["StudentID"])
	if err != nil {
		log.Fatal(err)
	}
	student.Score, err = strconv.ParseFloat(reply["Score"], 64)
	if err != nil {
		log.Fatal(err)
	}
	return student, err
}
