package main

import (
	"github.com/gin-gonic/gin"
	"Go_Lang_Api/db"
	"Go_Lang_Api/handlers"
)






// var books=[]book{
// 	{
// 		Id:       "1",
// 		Title:    "Go Basics",
// 		Author:   "Sahil",
// 		Quantity: 10,
// 	},
// 	{
// 		Id:       "2",
// 		Title:    "Learning APIs",
// 		Author:   "Rahul",
// 		Quantity: 5,
// 	},
// 	{
// 		Id:       "3",
// 		Title:    "Mastering Go",
// 		Author:   "Priya",
// 		Quantity: 7,
// 	},
// 	{
// 		Id:       "4",
// 		Title:    "Backend Development",
// 		Author:   "Aman",
// 		Quantity: 12,
// 	},
// }
// func getBookById(id string)(*book,error){
//    for i,b :=range books{
// 	  if b.Id==id{
// 		return &books[i],nil
// 	  }
// 	}
// 	  return nil,errors.New("book not found")
// }
// func bookById(c *gin.Context){
// 	id:=c.Param("id")
// 	book,err := getBookById(id)

// 	if err !=nil{
// 		c.IndentedJSON(http.StatusNotFound,gin.H{
// 			"message":err.Error(),
// 		})
// 		return
// 	}

// 	c.IndentedJSON(http.StatusOK,book)

// }

// func getBooks(c *gin.Context){   //Gin gives this object containing request + response details.
//         c.IndentedJSON(http.StatusOK,books)   //give indented json format of response you can take any
// }

// func createBook(c *gin.Context){
// 	var newBook book

// 	if err :=c.BindJSON(&newBook); err!=nil{
// 		return 
// 	}
// 	books=append(books,newBook)
// 	c.IndentedJSON(http.StatusCreated,newBook)
// }

// func checkoutbook(c *gin.Context){// query parameter
//    id,ok:=c.GetQuery("id")

//    if !ok{
// 	c.IndentedJSON(http.StatusBadRequest,gin.H{
// 		"message":"Missing Id query Paramter",
// 	})
//    }
//    book,err:=getBookById(id)

//    if err!=nil{
//        c.IndentedJSON(http.StatusNotFound,gin.H{
// 			"message":err.Error(),
// 		})
// 		return
//    }

//    if book.Quantity <=0{
// 	 c.IndentedJSON(http.StatusNotFound,gin.H{
// 			"message":"Book not found",
// 		})
// 		return
//    }

//     book.Quantity -=1
// 	c.IndentedJSON(http.StatusOK,book)
// }
//  func returnbook(c *gin.Context){
// 	id,ok:=c.GetQuery("id")

//    if !ok{
// 	c.IndentedJSON(http.StatusBadRequest,gin.H{
// 		"message":"Missing Id query Paramter",
// 	})
//    }
//    book,err:=getBookById(id)

//    if err!=nil{
//        c.IndentedJSON(http.StatusNotFound,gin.H{
// 			"message":err.Error(),
// 		})
// 		return
//    }

   

//     book.Quantity +=1
// 	c.IndentedJSON(http.StatusOK,book)
//  }
func main() {

	dbPool := db.ConnectDB()

	defer dbPool.Close()

	router := gin.Default()

	bookHandler := handlers.BookHandler{
		DB: dbPool,
	}

	router.GET(
		"/books",
		bookHandler.GetBooks,
	)

	router.POST(
		"/books",
		bookHandler.CreateBook,
	)

	router.Run(":8080")
}