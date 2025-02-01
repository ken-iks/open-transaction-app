package utils

import (
	"fmt"
	"log"
	"maps"
	"reflect"
	"strconv"
	"strings"
	"golang.org/x/exp/slices"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var testdb *gorm.DB

/*

initTestDB() initializes the database sqlite database - either starting
a new one, or opening one that already exists. 

I set automatic logging to silent because I do my own error loggin (redundacy)

*/
func initTestDB() {
	var err error
	testdb, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("Failed to connect to SQLite:", err)
	}
	fmt.Println("Connected to SQLite")

	// AutoMigrate models
	testdb.AutoMigrate(&Message{})
}

/*

Message struct type definition. Since we store 2 sets of routing and
account number pairs, I abstracted these into ClientInfo objects, then
just store one for the sender and one for the receiver

*/

type ClientInfo struct {
	Routing_num string `json:"routing_num"`
	Account_num string `json:"account_num"`
}

type Message struct {
	gorm.Model
	Seq int `json:"seq"`
	SenderInfo ClientInfo `gorm:"serializer:json" json:"sender_info"`
	ReceiverInfo ClientInfo `gorm:"serializer:json" json:"receiver_info"`
	Amount int `json:"amount"`
}

/*

Helper functions for Create_message function:

break_string(s) takes a string input that is supposed to be of
the form 'key=val'. The function returns each individual string
or a pair of empty strings if the input is invalid

Is_valid(s) takes a string input that is supposed to be a val field
from the message string. All value fields must only contain digits,
so function checks that the strings can be converted to an integer,
and returns true if it can

*/

func break_string(s string) (string, string) {
	split  := strings.Split(s, "=")
	if len(split) != 2 {
		fmt.Println("Incorrect format: Substring must be of the form key=value")
		return "",""
	}
	if len(split[0]) < 1 || len(split[1]) < 1 {
		fmt.Println("Incorrect format: Empty key or val field found in message")
		return "",""
	}
	return split[0], split[1]
}

func Is_valid(s string) (bool) {
	_, err := strconv.Atoi(s)
	return err == nil
}

/*

Create_message(input, user_acn, user_rtn) is the main function for instantiating a 
Message object to be added to our database. The input string goes through various
checks to make sure that it is in a valid format. If valid, return a Message object 
representing the input string. If not, log the appropriate error, and return an
empty Message object

*/

func Create_message(input_message string, user_acn string, user_rtn string) (Message) {
	// simple check to make sure string is not empty
	if len(input_message) < 10 {
		fmt.Println("Incorrect format: Must be 6 key=val pairs seperated by ; delimiter")
		return Message{}
	}
	// check string has 6 substrings seperate byt ';' delimiter
	message_arr := strings.Split(input_message, ";")
	if len(message_arr) != 6 {
		fmt.Println("Incorrect format: Must be 6 key=val pairs seperated by ; delimiter")
		return Message{}
	}

	// creating and populating a map of keys to values for each of the 6 substrings
	var m  = make(map[string]string, 6)

	for _, message := range message_arr {
		// check if substring is in correct 'key=val' format
		command, value := break_string(message)
		if command == "" || value == "" {
			return Message{}
		}
		// check that val field only contains integer characters
		if Is_valid(value) {
			m[command] = value
		} else {
			fmt.Println("Incorrect format: one of the value fields contains a non integer character")
			return Message{}
		}
	}
	// creating an array of only the keys in the map
	var keyarr [6]string
	var curr_ind = 0
	for val := range maps.Keys(m) {
		keyarr[curr_ind] = val
		curr_ind++
		}

	// defining a reference array of valid keys (the keys that should be in keyarr)
	var VALIDKEYS = [6]string{"seq", "sender_rtn", "sender_an", "receiver_rtn", "receiver_an", "amount"}
	
	// keys should be able to be input in any order
	slices.Sort(keyarr[:])
	slices.Sort(VALIDKEYS[:])
	
	// check that key set is equal to the valid keyset
	if !reflect.DeepEqual(keyarr, VALIDKEYS) {
		fmt.Println(("Incorrect format: Invalid KeySet"))
		return Message{}
	}

	// Light OAuth: Check if sender_an and sender_rtn match the current user
	if !(m["sender_an"]==user_acn && m["sender_rtn"]==user_rtn) {
		fmt.Println(("Incorrect Authorization: Sender Info Doesn't Match User Info"))
		fmt.Println("USER ACCOUNT NUM: ", user_acn)
		fmt.Println("USER ROUTING NUM: ", user_rtn)
		return Message{}
	}

	// Populate Message Struct
	senderInfo := ClientInfo{
		Routing_num: m["sender_rtn"],
		Account_num: m["sender_an"],
	}

	recieverInfo := ClientInfo{
		Routing_num: m["receiver_rtn"],
		Account_num: m["receiver_an"],
	}

	seq_id, seq_err := strconv.Atoi(m["seq"])

	// Check that sequence number is greater than 0
	if seq_err != nil || seq_id < 0 {
		fmt.Println("Incorrect format: Seq ID must be a positive integer")
		return Message{}
	}
	// Check that amount is greater than 0
	amount, amount_error := strconv.Atoi(m["amount"])
	if amount_error != nil || amount < 0 {
		fmt.Println("Incorrect format: Amount not an integer")
		return Message{}
	}

	return Message{
		Seq: seq_id,
		SenderInfo: senderInfo,
		ReceiverInfo: recieverInfo,
		Amount: amount,
	}
}

/*

Db_insert(msg) takes in a message struct and attempts to add it to the DB.
Only insert the message if it both a valid message (not empty) and has also
has a unique sequence id (to avoid repeats in DB)

*/

func DB_insert(message Message) {
	null_msg := Message{}
	if message == null_msg {
		fmt.Println("Cannot add invalid message to DB")
		return
	}

	initTestDB()
	// check if seq id already in DB
	var placeholder Message
	err := testdb.First(&placeholder, "seq = ?", message.Seq).Error
	if err == nil {
		fmt.Println("Seq ID already in DB")
		return
	} else if err != gorm.ErrRecordNotFound {
		fmt.Println("Error checking database:", err)
		return
	} else {
		// If all is good, add the message
		testdb.Create(&message)
		fmt.Println("Added to DB. Seq ID: ", message.Seq)
	}
}

/*

Db_fetch(seq) takes in a sequence number that you would like to fetch from the
database. Return an empty message list if the sequence number is not currently
in the database. If sequence number is less than zero, return the entire list
of Message rows in the database

*/

func DB_fetch(seq int) ([]Message){
	initTestDB()
	if seq < 0 {
		var all_messages []Message
		testdb.Find(&all_messages)
		return all_messages
	} else {
		var message Message
		result := testdb.First(&message, "seq = ?", seq) 
		if result.Error != nil {
			fmt.Println("Error fetching message: Seq ID not in DB")
			return []Message{}
		} else {
			return []Message{message}
		}
	}
}

/*

DeleteDB() just clears the database. Mainly for use in testing

*/

func DeleteDB() {
	initTestDB()
	testdb.Migrator().DropTable(&Message{}) // Drop the table
	testdb.AutoMigrate(&Message{})          // Recreate the table
}