package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

//---------------------------------show_tables-----------------------------------------
func show_tables() {

	var dbname string
	fmt.Println("Enter Database name to show its tables")
	fmt.Scanln(&dbname)
	param := "Manar:Mero144*@tcp(192.168.161.129:3306)/"
	db, err := sql.Open("mysql", param+dbname)
	res, _ := db.Query("SHOW TABLES")
	var table string
	for res.Next() {
		res.Scan(&table)
		fmt.Println(table)
	}
	fmt.Println(db, err)

}

//------------------------------------selectt---------------------------------------
func selectt() {
	fmt.Println("select tables you want to show ")

	var dbname string
	fmt.Println("Enter Database name")
	fmt.Scanln(&dbname)
	param := "Manar:Mero144*@tcp(192.168.161.129:3306)/"
	db, err := sql.Open("mysql", param+dbname)
	var table_name string
	fmt.Println("Enter table name")
	fmt.Scanln(&table_name)
	res, err := db.Query("select * from " + table_name)

	//------------select---------------user Table------------------------------
	if table_name == "user" {
		type user struct {
			user_ID           int
			firstName_user    string
			middleName_user   string
			lastName_user     string
			username          string
			mobile_user       string
			email_user        string
			password_user     string
			registeredAt_user string
			lastLogin_user    string
		}
		for res.Next() {
			var s user
			err = res.Scan(&s.user_ID, &s.firstName_user, &s.middleName_user, &s.lastName_user, &s.username, &s.mobile_user, &s.email_user, &s.password_user, &s.registeredAt_user, &s.lastLogin_user)
			println(s.user_ID, "  ", s.firstName_user, "  ", s.middleName_user, "  ", s.lastName_user, "  ", s.username, "  ", s.mobile_user, "  ", s.email_user, "  ", s.password_user, "  ", s.registeredAt_user, "  ", s.lastLogin_user)
		}
	}

	//----------------select------------user_follower--------------------------------
	if table_name == "user_follower" {
		type user_follower struct {
			user_follower_ID       int
			type_userFollower      string
			createdAt_userFollower string
			updatedAt_userFollower string
			user_ID                int
		}
		for res.Next() {
			var s user_follower
			err = res.Scan(&s.user_follower_ID, &s.user_follower_ID, &s.createdAt_userFollower, &s.updatedAt_userFollower, &s.user_ID)
			println(s.user_follower_ID, "   ", s.type_userFollower, "   ", s.createdAt_userFollower, "   ", s.updatedAt_userFollower, s.user_ID)
		}
	}
	//-------------select--------------------user_post------------------------------------
	if table_name == "user_post" {
		type user_post struct {
			user_post_ID       int
			post_content       string
			post_language      string
			createdAt_userPost string
			updatedAt_userPost string
			user_ID            int
		}
		for res.Next() {
			var s user_post
			err = res.Scan(&s.user_post_ID, &s.post_content, &s.post_language, &s.createdAt_userPost, &s.updatedAt_userPost, &s.user_ID)
			println(s.user_post_ID, "   ", s.post_content, "   ", s.post_language, "   ", s.createdAt_userPost, "  ", s.updatedAt_userPost, "  ", s.user_ID)
		}
	}
	//---------------------select--------------------_group Table------------------------------------
	if table_name == "_group" {
		type _group struct {
			group_ID        int
			createdBy_group int
			updatedBy_group int
			createdAt_group string
			updatedAt_group string
			content_group   string
			status_group    string
			user_ID         int
		}
		for res.Next() {
			var s _group
			err = res.Scan(&s.group_ID, &s.createdBy_group, &s.updatedBy_group, &s.createdAt_group, &s.updatedAt_group, &s.content_group, &s.status_group)
			println(s.group_ID, "   ", s.createdBy_group, "   ", s.updatedBy_group, "   ", s.createdAt_group, "   ", s.updatedAt_group, "   ", s.content_group, "   ", s.status_group)
		}
	}
	//---------------------select--------------------group_message Table------------------------------------
	if table_name == "group_message" {
		type group_message struct {
			groupMessage_ID        int
			message_body           string
			createdAt_groupMessage string
			updatedAt_groupMessage string
			group_ID               int
			user_ID                int
		}
		for res.Next() {
			var s group_message
			err = res.Scan(&s.groupMessage_ID, &s.message_body, &s.createdAt_groupMessage, &s.updatedAt_groupMessage, &s.group_ID, &s.user_ID)
			println(s.groupMessage_ID, "   ", s.message_body, "   ", s.createdAt_groupMessage, "   ", s.updatedAt_groupMessage, s.group_ID, &s.user_ID)
		}
	}
	//---------------------select--------------------group_meta Table------------------------------------
	if table_name == "group_meta" {
		type group_meta struct {
			groupMeta_ID      int
			groupMeta_content string
			group_ID          int
		}
		for res.Next() {
			var s group_meta
			err = res.Scan(&s.groupMeta_ID, &s.groupMeta_content, &s.group_ID)
			println(s.groupMeta_ID, "   ", s.groupMeta_content, "   ", s.group_ID, "   ")
		}
	}

	fmt.Println(db, err)

}

//-------------------------------------------------insert fun-------------------------------------------------
func insert() {
	fmt.Println("insert choice here")
	var dbname string
	fmt.Println("Enter Database name")
	fmt.Scanln(&dbname)
	param := "Manar:Mero144*@tcp(192.168.161.129:3306)/"
	db, err := sql.Open("mysql", param+dbname)
	var table_name string
	fmt.Println("Enter table name")
	fmt.Scanln(&table_name)
	//---------------------insert--------------------user Table------------------------------------
	if table_name == "user" {

		var user_ID int
		var firstName_user string
		var middleName_user string
		var lastName_user string
		var username string
		var mobile_user string
		var email_user string
		var password_user string
		var registeredAt_user string
		var lastLogin_user string

		println("Enter Data:")
		print("user id:  ")
		fmt.Scanln(&user_ID)
		print("firstName_user:  ")
		fmt.Scanln(&firstName_user)
		print("middleName_user:  ")
		fmt.Scanln(&middleName_user)
		print("lastName_user:  ")
		fmt.Scanln(&lastName_user)
		print("username:  ")
		fmt.Scanln(&username)
		print("mobile_user:  ")
		fmt.Scanln(&mobile_user)
		print("email_user:  ")
		fmt.Scanln(&email_user)
		print("password_user:  ")
		fmt.Scanln(&password_user)
		print("registeredAt_user:  ")
		fmt.Scanln(&registeredAt_user)
		print("lastLogin_user:  ")
		fmt.Scanln(&lastLogin_user)

		db.Query("insert into " + table_name + " values(" + strconv.Itoa(user_ID) + ",'" + firstName_user + "','" + middleName_user + "','" + lastName_user + "','" + username + "','" + mobile_user + "','" + email_user + "','" + password_user + "', '" + registeredAt_user + "', '" + lastLogin_user + "')")
	}
	//---------------------insert--------------------user_follower Table------------------------------------
	if table_name == "user_follower" {

		var user_follower_ID int
		var type_userFollower string
		var createdAt_userFollower string
		var updatedAt_userFollower string
		var user_ID int

		println("Enter Data:")
		print("user_follower_ID:  ")
		fmt.Scanln(&user_follower_ID)
		print("type_userFollower: ")
		fmt.Scanln(&type_userFollower)
		print("createdAt_userFollower:  ")
		fmt.Scanln(&createdAt_userFollower)
		print("updatedAt_userFollower:  ")
		fmt.Scanln(&updatedAt_userFollower)
		print("user_ID:  ")
		fmt.Scanln(&user_ID)

		db.Query("insert into " + table_name + " values(" + strconv.Itoa(user_follower_ID) + ",'" + type_userFollower + "','" + createdAt_userFollower + "','" + updatedAt_userFollower + "','" + strconv.Itoa(user_ID) + "')")
	}
	//---------------------insert--------------------user_post Table------------------------------------
	if table_name == "user_post" {

		var user_post_ID int
		var post_content string
		var post_language string
		var createdAt_userPost string
		var updatedAt_userPost string
		var user_ID int

		println("Enter Student Data:")
		print("user_post_ID:  ")
		fmt.Scanln(&user_post_ID)
		print("post_content: ")
		fmt.Scanln(&post_content)
		print("post_language:  ")
		fmt.Scanln(&post_language)
		print("createdAt_userPost: ")
		fmt.Scanln(&createdAt_userPost)
		print("updatedAt_userPost: ")
		fmt.Scanln(&updatedAt_userPost)
		print("user_ID: ")
		fmt.Scanln(&user_ID)
		db.Query("insert into " + table_name + " values(" + strconv.Itoa(user_post_ID) + ",'" + post_content + "','" + post_language + "','" + createdAt_userPost + "','" + "','" + updatedAt_userPost + "','" + strconv.Itoa(user_ID) + "')")
	}

	fmt.Println(db, err)
}

//----------------------------------update-------------------------
func update() {
	fmt.Println("update choice here")
	db, err := sql.Open("mysql", "Manar:Mero144*@tcp(192.168.161.129:3306)/Facebook")

	var table_name string
	fmt.Println("Enter table name")
	fmt.Scanln(&table_name)
	//---------------------update--------------------_group Table------------------------------------
	if table_name == "_group" {
		var group_ID int
		var content_group string
		var status_group string

		println("Enter ID to search:")
		print("group_ID: ")
		fmt.Scanln(&group_ID)

		print("content_group  ")
		fmt.Scanln(&content_group)

		print("status_group: ")
		fmt.Scanln(&status_group)

		db.Query("UPDATE " + table_name + " SET group_ID ='" + strconv.Itoa(group_ID) + "' WHERE group_ID=" + strconv.Itoa(group_ID) + "")
		db.Query("UPDATE " + table_name + " SET content_group ='" + content_group + "' WHERE group_ID=" + strconv.Itoa(group_ID) + "")
		db.Query("UPDATE " + table_name + " SET status_group ='" + status_group + "' WHERE group_ID=" + strconv.Itoa(group_ID) + "")

	}
	//---------------------update--------------------group_meta Table------------------------------------
	if table_name == "group_meta" {

		var groupMeta_ID int
		var groupMeta_content string
		var group_ID int

		println("Enter ID to search:")
		print("G ID:  ")
		fmt.Scanln(&groupMeta_ID)
		print("Group Meta Content: ")
		fmt.Scanln(&groupMeta_content)
		print("group_ID :  ")
		fmt.Scanln(&group_ID)

		db.Query("UPDATE " + table_name + " SET groupMeta_ID ='" + strconv.Itoa(groupMeta_ID) + "' WHERE groupMeta_ID=" + strconv.Itoa(groupMeta_ID) + "")
		db.Query("UPDATE " + table_name + " SET groupMeta_content ='" + groupMeta_content + "' WHERE groupMeta_ID=" + strconv.Itoa(groupMeta_ID) + "")
		db.Query("UPDATE " + table_name + " SET group_ID ='" + strconv.Itoa(group_ID) + "' WHERE groupMeta_ID=" + strconv.Itoa(groupMeta_ID) + "")

	}

	//---------------------update--------------------group_post Table------------------------------------
	if table_name == "group_post" {
		var groupPost_ID int
		var createdAt_groupPost string
		var groupPost_content string

		println("Enter ID to search:")
		print("groupPost_ID: ")
		fmt.Scanln(&groupPost_ID)

		print("createdAt_groupPost  ")
		fmt.Scanln(&createdAt_groupPost)

		print("status_group: ")
		fmt.Scanln(&groupPost_content)

		db.Query("UPDATE " + table_name + " SET groupPost_ID ='" + strconv.Itoa(groupPost_ID) + "' WHERE groupPost_ID=" + strconv.Itoa(groupPost_ID) + "")
		db.Query("UPDATE " + table_name + " SET createdAt_groupPost ='" + createdAt_groupPost + "' WHERE groupPost_ID=" + strconv.Itoa(groupPost_ID) + "")
		db.Query("UPDATE " + table_name + " SET createdAt_groupPost ='" + createdAt_groupPost + "' groupPost_ID=" + strconv.Itoa(groupPost_ID) + "")

	}
	println(err, db, "done")

}

//---------------------------delete_from_table_by_id----------------------
func delete_from_table_by_id() {
	db, err := sql.Open("mysql", "Manar:Mero144*@tcp(192.168.161.129:3306)/Facebook")
	var table_name string
	fmt.Println("Enter table name")
	fmt.Scanln(&table_name)

	//---------------------delete--------------------user_follower Table------------------------------------
	if table_name == "user_follower" {
		var user_follower_ID int
		print("Enter ID: ")
		fmt.Scanln(&user_follower_ID)
		db.Query("delete from " + table_name + " where user_follower_ID='" + strconv.Itoa(user_follower_ID) + "'")
		fmt.Println(db, err)

	}
	//---------------------delete--------------------group_meta Table------------------------------------
	if table_name == "group_meta" {
		var groupMeta_ID int
		print("Enter ID: ")
		fmt.Scanln(&groupMeta_ID)
		db.Query("delete from " + table_name + " where groupMeta_ID='" + strconv.Itoa(groupMeta_ID) + "'")
		fmt.Println(db, err)

	}
	//---------------------delete--------------------user_message Table------------------------------------
	if table_name == "user_message" {
		var user_message_ID int

		print("ID_groupMessage: ")
		fmt.Scanln(&user_message_ID)
		db.Query("delete from " + table_name + " where user_message_ID='" + strconv.Itoa(user_message_ID) + "'")
		fmt.Println(db, err)

	}

}

func main() {

	db, err := sql.Open("mysql", "Salman:MS123123@tcp(192.168.161.129:3306)/Facebook")
	fmt.Println(db, err)

	fmt.Println("main menu")
	fmt.Println("Press 1 to Show Tables")
	fmt.Println("Press 2 to Select a Table")
	fmt.Println("Press 3 to insert")
	fmt.Println("Press 4 to update")
	fmt.Println("Press 5 to Delete by id")
	fmt.Println("Press 6 to exit")

	for {
		choice()
	}
}

func choice() {
	fmt.Println("Enter your valid number")
	var input int
	n, err := fmt.Scanln(&input)
	if n < 1 || err != nil {
		fmt.Println("invalid input")

	}

	switch input {
	case 1:
		show_tables()
		break
	case 2:
		selectt()
		break
	case 3:
		insert()
		break
	case 4:
		update()
		break
	case 5:
		delete_from_table_by_id()
		break
	case 6:
		os.Exit(11)
		break
	default:
		fmt.Println("Facebook Database")
	}
}
