package main

import (
	"backend/tests"
	"backend/utils"
	"fmt"
	"net/http"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"

	"github.com/gin-gonic/gin"
)

/*

LoginRequest struct is equivalent to the ClientInfo strruction, 
just used to store the information of the currently logged in
user. These are used in the Create_message function as an 
authentication step.

*/
type LoginRequest struct {
	AccountNumber string `json:"account_number"`
	RoutingNumber string `json:"routing_number"`
}

/*

testSuite() runs the internal tests for testing some of the edge cases
associated with creating a message and storing it in the database.

*/
func testSuite() {
	tests.Create_message_tests()
}

/*

get_messages(c) uses the Db_fetch function with 'id' which is just the sequence
number of the message you are looking for. This sends the message list generated
by the function back to the frontend to be displayed in the table.

*/
func get_messages(c *gin.Context) {
	var requestData struct {
		Input int `json:"id"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request", "details": err.Error()})
		return
	}
	id := requestData.Input

	message_list := utils.DB_fetch(id)
	fmt.Printf("%+v\n", message_list)

	c.JSON(200, message_list)
}

/*

add_message_toDB(c) uses the Create_message and DB_insert function to attempt
to add a message string to our database. Returns error if the message has an invalid
format. If message has a valid format but DB doesn't change, then it means sequence 
number is already in the DB.

*/

// Add message to DB
func add_message_toDB(c *gin.Context) {
	var requestData struct {
		Input string `json:"input"`
	}
	if err := c.ShouldBindJSON(&requestData); err != nil {
		c.JSON(400, gin.H{"message": "Invalid request"})
		return
	}
	
	fmt.Printf("\nBACKEND RECIEVED STRING: %s. ATTEMPTING TO ADD TO DATABASE\n\n", requestData.Input)
	session := sessions.Default(c)
	user_acn := session.Get("account_number")
	user_rtn := session.Get("routing_number")

	fmt.Printf("Session Data - Account Number: %v, Routing Number: %v\n", user_acn, user_rtn)

	// convert user info to string
	user_acn_string, _ := user_acn.(string)
	user_rtn_string, _ := user_rtn.(string)
	message_struct := utils.Create_message(requestData.Input, user_acn_string , user_rtn_string)

	// check if valid Message object is returned from Create_message function
	null_msg := utils.Message{}
	if message_struct == null_msg {
		responseMessage := ("Unable to add to Database. Please check format of message")
		c.JSON(400, gin.H{"message": responseMessage})
		return
	}
	// if valid, attempt to insert in DB - check logs for any potential errors
	utils.DB_insert(message_struct)
	responseMessage := ("Valid message format! Added to DB")
	c.JSON(200, gin.H{"message": responseMessage})
}


func main() {
	router := gin.Default()

	// Defining permissions
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// Handle preflight requests
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("session", store))

	// LOGIN ENPOINT

	router.POST("/login", func(c *gin.Context) {
		var loginData LoginRequest
		if err := c.ShouldBindJSON(&loginData); err != nil {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		fmt.Printf("User Attempting Login: Account: %s, Routing: %s\n", 
		loginData.AccountNumber, loginData.RoutingNumber)

		if !(utils.Is_valid(loginData.AccountNumber) || utils.Is_valid(loginData.RoutingNumber)) {
			c.JSON(400, gin.H{"error": "Invalid request"})
			return
		}

		session := sessions.Default(c)
		session.Set("account_number", loginData.AccountNumber)
		session.Set("routing_number", loginData.RoutingNumber)
		session.Save()

		c.JSON(200, gin.H{"message": "Login successful"})
	
	})

	// LOGOUT ENDPOINT

	router.POST("/logout", func(c *gin.Context) {
		session := sessions.Default(c)
		session.Clear()
		session.Save()
		c.JSON(200, gin.H{"message": "Logged out successfully"})
	})

	// SESSION ENPOINT (check if a user is logged in)

	router.GET("/session", func(c *gin.Context) {
		session := sessions.Default(c)
		account := session.Get("account_number")
		if account == nil {
			c.JSON(401, gin.H{"error": "Not logged in"})
			return
		}
		c.JSON(200, gin.H{"message": "User is logged in", "account_number": account})
	})

	// Fetch messages
	router.POST("/messages", get_messages)

	// Add message to DB
	router.POST("/echo", add_message_toDB)

	//testSuite() -> was using for testing as I was building

	router.Run(":8080")
}