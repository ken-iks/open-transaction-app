package tests

import (
	"fmt"
	"backend/utils"
)

func Create_message_tests() {
	
	fmt.Printf("\nCLEARING TABLE\n")
	utils.DeleteDB()

	fmt.Printf("\nTEST CREATION + INSERTION 1\n\n")
	var test_regular = utils.Create_message("seq=1;sender_rtn=021000021;sender_an=537646894897833;receiver_rtn=121145307;receiver_an=669907820975207;amount=3424",
	"537646894897833", "021000021")
	utils.DB_insert(test_regular)
	
	fmt.Printf("\nTEST CREATION + INSERTION 2\n\n")
	var test_regular2 = utils.Create_message("seq=2;sender_rtn=121000248;sender_an=349848983426759;receiver_rtn=121145307;receiver_an=160661577716921;amount=2123",
	"349848983426759","121000248")
	utils.DB_insert(test_regular2)
	
	fmt.Printf("\nTEST CREATION + INSERTION OF EMPTY STRING\n\n")
	var test_empty_str = utils.Create_message("", "", "")
	utils.DB_insert(test_empty_str)
	
	fmt.Printf("\nTEST CREATION + INSERTION OF INVALID AMOUNT\n\n")
	var test_invalid_amount = utils.Create_message("seq=1;sender_rtn=021000021;sender_an=537646894897833;receiver_rtn=121145307;receiver_an=669907820975207;amount=nah",
	"537646894897833", "021000021")
	utils.DB_insert(test_invalid_amount)

	fmt.Printf("\nTEST CREATION + INSERTION OF INVALID SEQ\n\n")
	var test_invalid_seq = utils.Create_message("seq=-1;sender_rtn=021000021;sender_an=537646894897833;receiver_rtn=121145307;receiver_an=669907820975207;amount=43242",
	"537646894897833", "021000021")
	utils.DB_insert(test_invalid_seq)

	fmt.Printf("\nTEST CREATION + INSERTION OF MESSAGE ALREADY IN DB\n\n")
	var test_repeat = utils.Create_message("seq=2;sender_rtn=121000248;sender_an=349848983426759;receiver_rtn=121145307;receiver_an=160661577716921;amount=2123",
	"349848983426759", "121000248")
	utils.DB_insert(test_repeat)

	fmt.Printf("\nTEST CREATION + INSERTION OF MESSAGE WITH NEGATIVE SEQ\n\n")
	var test_negative_seq = utils.Create_message("seq=-5;sender_rtn=121000248;sender_an=349848983426759;receiver_rtn=121145307;receiver_an=160661577716921;amount=2123",
	"349848983426759", "121000248")
	utils.DB_insert(test_negative_seq)

	fmt.Printf("\nTEST CREATION + INSERTION OF MESSAGE WITH AN EMPTY VALUE FIELD\n\n")
	var test_empty_field = utils.Create_message("seq=6;sender_rtn=;sender_an=349848983426759;receiver_rtn=121145307;receiver_an=160661577716921;amount=2123",
	"349848983426759", "")
	utils.DB_insert(test_empty_field)

	fmt.Printf("\nTEST CREATION + INSERTION OF MESSAGE WITH AN INVALID VALUE FIELD\n\n")
	var test_invalid_field = utils.Create_message("seq=6;sender_rtn=1210j00248;sender_an=349848983426759;receiver_rtn=121145307;receiver_an=160661577716921;amount=2123",
	"349848983426759","1210j00248")
	utils.DB_insert(test_invalid_field)

	fmt.Printf("\nTEST CREATION + INSERTION OF MESSAGE WITH AN INVALID AUTHORIZATION (sender != user info)\n\n")
	var test_invalid_login = utils.Create_message("seq=10;sender_rtn=021000021;sender_an=629385443170308;receiver_rtn=121145307;receiver_an=136657407199052;amount=10343",
	"10","11")
	utils.DB_insert(test_invalid_login)
	
	fmt.Printf("\nTEST FETCH ALL ROWS FROM TABLE\n\n")
	var all_db_rows = utils.DB_fetch(-1)
	fmt.Printf("Full DB:\n %+v\n", all_db_rows)

	fmt.Printf("\nTEST FETCH VALID ROW IN TABLE\n\n")
	var specific_db_row = utils.DB_fetch(2)
	fmt.Printf("Specific DB row:\n %+v\n", specific_db_row)

	fmt.Printf("\nTEST FETCH INVALID ROW FROM TABLE\n\n")
	var test_erronious_db_row = utils.DB_fetch(20)
	fmt.Printf("Erronious DB row:\n %+v\n", test_erronious_db_row)

	fmt.Printf("\nCLEARING TABLE\n")
	utils.DeleteDB()

}