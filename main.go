package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB
var err error

type Person struct{
    ID uint `json:"id"`
    FirstName string `json:"firstname"`
    LastName string `js9on:"lastname"`
}
   func main() {

    db, err = gorm.Open("sqlite3", "./gorm.db")
    if err != nil {
        fmt.Println(err)
      }
    defer db.Close()

    db.AutoMigrate(&Person{})

    r := gin.Default()
    r.GET("/people", GetPeople)
    r.GET("/people/:id",GetPerson)
      r.POST("/people",CreatePerson)
    r.DELETE("/people/:id",DeletePerson)

    r.Run(":8080")

    // p1 := Person{FirstName: "Bibek", LastName: "Magar"}
    // p2 := Person{FirstName: "Ram", LastName: "Shrestha"}
    // fmt.Println(p1.FirstName)
    // fmt.Println(p2.LastName)

    // db.Create(&p1)
    // var p3 Person
    // db.First(&p3)

    // fmt.Println(p1.FirstName)
    // fmt.Println(p2.LastName)
    // fmt.Println(p3.LastName)
   }

   func CreatePerson(c *gin.Context){
       var person Person
       c.BindJSON(&person)

       db.Create(&person)
       c.JSON(200,person)
   }
   
   func DeletePerson(c *gin.Context) {
    id := c.Params.ByName("id")
    var person Person
    d := db.Where("id = ?", id).Delete(&person)
    fmt.Println(d)
    c.JSON(200, gin.H{"id #" + id: "deleted"})
   }

   func GetPeople(c *gin.Context){
       var people []Person
       if err := db.Find(&people).Error; err != nil{
           c.AbortWithStatus(404)
           fmt.Println(err)
       }else{
           c.JSON(200,people)
       }
   }

   func GetPerson(c *gin.Context){
       fmt.Println("Hello")
    id := c.Params.ByName("id")
       var person Person
       if err := db.Where("id = ?",id).First(&person).Error; err !=nil{
           c.AbortWithStatus(404)
           fmt.Println(err)
       }else {
           c.JSON(200,person)
       }
   }