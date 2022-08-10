package main

import (
	"database/sql"
	"fmt"
	"image/color"
	"os"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"io/ioutil"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "Manar:Mero144*@tcp(192.168.43.74:3306)/Facebook")
	fmt.Println(db, err)

	a := app.New()
	w := a.NewWindow("Facebook Database")
	w.Resize(fyne.NewSize(700, 700))

	result := canvas.NewText("", color.Black)
	result.TextSize = 80
	result.Alignment = fyne.TextAlignCenter

	entry1 := widget.NewEntry()
	entry1.TextStyle = fyne.TextStyle{Bold: true}
	entry1.SetPlaceHolder("Enter Database name    ")

	entry2 := widget.NewEntry()
	entry2.TextStyle = fyne.TextStyle{Bold: true}
	entry2.SetPlaceHolder("Enter Table name      ")

	//---------------------------------show_tables-----------------------------------------
	btn1 := widget.NewButton("Show tables", func() {
		err := os.Remove("C:/Users/LapTop MarKet/Desktop/data.txt")
		f, err := os.Create("C:/Users/LapTop MarKet/Desktop/data.txt")
		var dbname string
		dbname = entry1.Text
		param := "Manar:Mero144*@tcp(192.168.43.74:3306)/"
		db, err := sql.Open("mysql", param+dbname)
		res, _ := db.Query("SHOW TABLES")
		var table string
		for res.Next() {
			res.Scan(&table)
			_, err = f.WriteString(table + "\n" + "\n")
		}

		data, _ := ioutil.ReadFile("C:/Users/LapTop MarKet/Desktop/data.txt")

		result := fyne.NewStaticResource("All Tables", data)
		entry := widget.NewMultiLineEntry()
		entry.SetText(string(result.StaticContent))
		w := fyne.CurrentApp().NewWindow(
			string(result.StaticName))
		w.SetContent(container.NewScroll(entry))
		w.Resize(fyne.NewSize(400, 400))
		w.Show()

		fmt.Println(db, err)

	})

	//---------------------------------show_tables-----------------------------------------
	btn2 := widget.NewButton("Select a specific table to show all data", func() {

		err := os.Remove("C:/Users/LapTop MarKet/Desktop/data.txt")
		f, err := os.Create("C:/Users/LapTop MarKet/Desktop/data.txt")
		var dbname string
		dbname = entry1.Text
		param := "Manar:Mero144*@tcp(192.168.43.74:3306)/"
		db, err := sql.Open("mysql", param+dbname)
		res, err := db.Query("select * from " + entry2.Text)
		var all string
		//------------select---------------user Table------------------------------
		if entry2.Text == "user" {
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
				all = strconv.Itoa(s.user_ID) + "   " + s.firstName_user + "   " + s.middleName_user + "   " + s.lastName_user + "   " + s.username + "   " + s.mobile_user + "   " + s.email_user + "   " + s.password_user + "   " + s.registeredAt_user + "   " + s.lastLogin_user
				_, err = f.WriteString(all + "\n" + "\n")
			}
			data, _ := ioutil.ReadFile("C:/Users/LapTop MarKet/Desktop/data.txt")
			result := fyne.NewStaticResource("Data in Tables", data)
			entry := widget.NewMultiLineEntry()
			entry.SetText(string(result.StaticContent))
			w := fyne.CurrentApp().NewWindow(
				string(result.StaticName))
			w.SetContent(container.NewScroll(entry))
			w.Resize(fyne.NewSize(1000, 1000))
			w.Show()
		}
		//----------------select------------user_follower--------------------------------
		if entry2.Text == "user_follower" {
			type user_follower struct {
				user_follower_ID       int
				type_userFollower      string
				createdAt_userFollower string
				updatedAt_userFollower string
				user_ID                int
			}
			for res.Next() {
				var s user_follower
				err = res.Scan(&s.user_follower_ID, &s.type_userFollower, &s.createdAt_userFollower, &s.updatedAt_userFollower, &s.user_ID)
				all = strconv.Itoa(s.user_follower_ID) + "   " + s.type_userFollower + "   " + s.createdAt_userFollower + "   " + s.updatedAt_userFollower + "   " + strconv.Itoa(s.user_ID)
				_, err = f.WriteString(all + "\n" + "\n")
			}
			data, _ := ioutil.ReadFile("C:/Users/LapTop MarKet/Desktop/data.txt")
			result := fyne.NewStaticResource("Data in Tables", data)
			entry := widget.NewMultiLineEntry()
			entry.SetText(string(result.StaticContent))
			w := fyne.CurrentApp().NewWindow(
				string(result.StaticName))
			w.SetContent(container.NewScroll(entry))
			w.Resize(fyne.NewSize(1000, 1000))
			w.Show()

		}
		//-------------select--------------------user_post------------------------------------
		if entry2.Text == "user_post" {
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
				all = strconv.Itoa(s.user_post_ID) + "   " + s.post_content + "   " + s.post_language + "   " + s.createdAt_userPost + "  " + s.updatedAt_userPost + "  " + strconv.Itoa(s.user_ID)
				_, err = f.WriteString(all + "\n" + "\n")
			}
			data, _ := ioutil.ReadFile("C:/Users/LapTop MarKet/Desktop/data.txt")
			result := fyne.NewStaticResource("Data in Tables", data)
			entry := widget.NewMultiLineEntry()
			entry.SetText(string(result.StaticContent))
			w := fyne.CurrentApp().NewWindow(
				string(result.StaticName))
			w.SetContent(container.NewScroll(entry))
			w.Resize(fyne.NewSize(1000, 1000))
			w.Show()
		}

		fmt.Println(db, err)
	})

	//-------------------------------------------------insert into specific table-------------------------------------------------
	btn3 := widget.NewButton("insert", func() {
		db, err := sql.Open("mysql", "Manar:Mero144*@tcp(192.168.43.74:3306)/Facebook")
		//---------------------insert--------------------user_friend ------------------------------------
		if entry2.Text == "user_friend" {
			w := a.NewWindow("user_friend")
			w.Resize(fyne.NewSize(400, 400))

			user_friend_ID := widget.NewEntry()
			user_friend_ID.SetPlaceHolder("Enter user_friend_ID...")

			status_userFriend := widget.NewEntry()
			status_userFriend.SetPlaceHolder("Enter status_userFriend...")

			type_userFriend := widget.NewEntry()
			type_userFriend.SetPlaceHolder("Enter type_userFriend...")

			notes := widget.NewEntry()
			notes.SetPlaceHolder("Enter notes...")

			createdAt_userFriend := widget.NewEntry()
			createdAt_userFriend.SetPlaceHolder("Enter createdAt_userFriend...")

			updatedAt_userFriend := widget.NewEntry()
			updatedAt_userFriend.SetPlaceHolder("Enter updatedAt_userFriend...")

			user_ID := widget.NewEntry()
			user_ID.SetPlaceHolder("Enter user_ID...")

			submit_btn := widget.NewButton("insert", func() {

				var n_ID int
				var n_status string
				var n_typee string
				var n_notes string
				var n_createdAt string
				var n_updatedAt string
				var n_user_ID int

				n_ID, err = strconv.Atoi(user_friend_ID.Text)
				fmt.Print(err)
				n_status = status_userFriend.Text
				n_typee = type_userFriend.Text
				n_notes = notes.Text
				n_createdAt = createdAt_userFriend.Text
				n_updatedAt = updatedAt_userFriend.Text
				n_user_ID, err = strconv.Atoi(user_ID.Text)
				fmt.Print(err)

				db.Query("insert into " + entry2.Text + " values(" + strconv.Itoa(n_ID) + ",'" + n_status + "','" + n_typee + "','" + n_notes + "','" + n_createdAt + "','" + n_updatedAt + "','" + strconv.Itoa(n_user_ID) + "')")

				user_friend_ID.Text = ""
				status_userFriend.Text = ""
				type_userFriend.Text = ""
				notes.Text = ""
				createdAt_userFriend.Text = ""
				updatedAt_userFriend.Text = ""
				user_ID.Text = ""

				user_friend_ID.Refresh()
				status_userFriend.Refresh()
				type_userFriend.Refresh()
				notes.Refresh()
				createdAt_userFriend.Refresh()
				updatedAt_userFriend.Refresh()
				user_ID.Refresh()

			})

			w.SetContent(
				container.NewVBox(user_friend_ID, status_userFriend, type_userFriend, notes, createdAt_userFriend, updatedAt_userFriend, user_ID, submit_btn),
			)
			w.Show()

		}
	})

	//----------------------------------update-------------------------
	btn4 := widget.NewButton("Update", func() {
		db, err := sql.Open("mysql", "Manar:Mero144*@tcp(192.168.43.74:3306)/Facebook")

		//---------------------update--------------------_group Table------------------------------------
		if entry2.Text == "_group" {
			w := a.NewWindow("_group")
			w.Resize(fyne.NewSize(400, 400))
			group_ID := widget.NewEntry()
			group_ID.SetPlaceHolder("Enter group_ID ")
			content_group := widget.NewEntry()
			content_group.SetPlaceHolder("Enter content_group ")

			status_group := widget.NewEntry()
			status_group.SetPlaceHolder("Enter status_group ")

			submit_btn := widget.NewButton("Update", func() {

				var ID int
				var n_content_group string
				var n_status_group string

				ID, err = strconv.Atoi(group_ID.Text)
				fmt.Print(err)
				n_content_group = content_group.Text
				n_status_group = status_group.Text

				db.Query("UPDATE " + entry2.Text + " SET content_group ='" + n_content_group + "' WHERE group_ID=" + strconv.Itoa(ID) + "")
				db.Query("UPDATE " + entry2.Text + " SET status_group ='" + n_status_group + "' WHERE group_ID=" + strconv.Itoa(ID) + "")

				content_group.Text = ""
				status_group.Text = ""

				group_ID.Text = ""
				content_group.Refresh()
				status_group.Refresh()
				group_ID.Refresh()

			})

			w.SetContent(
				container.NewVBox(group_ID, content_group, status_group, submit_btn),
			)
			w.Show()

		}
		//---------------------update--------------------group_meta Table------------------------------------
		if entry2.Text == "group_meta" {

			w := a.NewWindow("group_meta")
			w.Resize(fyne.NewSize(400, 400))
			groupMeta_ID := widget.NewEntry()
			groupMeta_ID.SetPlaceHolder("Enter groupMeta_ID   ")
			groupMeta_content := widget.NewEntry()
			groupMeta_content.SetPlaceHolder("Enter groupMeta_content   ")

			submit_btn := widget.NewButton("Update", func() {

				var ID int
				var n_groupMeta_content string

				ID, err = strconv.Atoi(groupMeta_ID.Text)
				fmt.Print(err)
				n_groupMeta_content = groupMeta_content.Text

				db.Query("UPDATE " + entry2.Text + " SET groupMeta_content ='" + n_groupMeta_content + "' WHERE groupMeta_ID=" + strconv.Itoa(ID) + "")

				groupMeta_content.Text = ""
				groupMeta_ID.Text = ""

				groupMeta_content.Refresh()
				groupMeta_ID.Refresh()

			})

			w.SetContent(
				container.NewVBox(groupMeta_ID, groupMeta_content, submit_btn),
			)
			w.Show()

		}
		//---------------------update--------------------group_post Table------------------------------------
		if entry2.Text == "group_post" {

			w := a.NewWindow("group_post")
			w.Resize(fyne.NewSize(400, 400))
			groupPost_ID := widget.NewEntry()
			groupPost_ID.SetPlaceHolder("Enter groupPost_ID   ")
			createdAt_groupPost := widget.NewEntry()
			createdAt_groupPost.SetPlaceHolder("Enter createdAt_groupPost   ")
			groupPost_content := widget.NewEntry()
			groupPost_content.SetPlaceHolder("Enter groupPost_content   ")

			submit_btn := widget.NewButton("Update", func() {

				var ID int
				var n_createdAt_groupPost string
				var n_groupPost_content string

				ID, err = strconv.Atoi(groupPost_ID.Text)
				fmt.Print(err)
				n_createdAt_groupPost = createdAt_groupPost.Text
				n_groupPost_content = groupPost_content.Text

				db.Query("UPDATE " + entry2.Text + " SET n_createdAt_groupPost ='" + n_createdAt_groupPost + "' WHERE groupPost_ID=" + strconv.Itoa(ID) + "")
				db.Query("UPDATE " + entry2.Text + " SET groupPost_content ='" + n_groupPost_content + "' WHERE groupPost_ID=" + strconv.Itoa(ID) + "")

				groupPost_ID.Text = ""
				createdAt_groupPost.Text = ""
				groupPost_content.Text = ""

				groupPost_ID.Refresh()
				createdAt_groupPost.Refresh()
				groupPost_content.Refresh()

			})

			w.SetContent(
				container.NewVBox(groupPost_ID, createdAt_groupPost, groupPost_content, submit_btn),
			)
			w.Show()

		}

	})

	//---------------------------Delete_row_from_table_by_id----------------------
	btn5 := widget.NewButton("Delete", func() {
		db, err := sql.Open("mysql", "Manar:Mero144*@tcp(192.168.43.74:3306)/Facebook")

		//---------------------delete--------------------group_meta Table------------------------------------
		if entry2.Text == "group_meta" {

			w := a.NewWindow("group_meta")
			w.Resize(fyne.NewSize(400, 400))
			id := widget.NewEntry()
			id.SetPlaceHolder("Enter group_meta_ID...")

			submit_btn := widget.NewButton("Delete", func() {

				var ID int
				ID, err = strconv.Atoi(id.Text)
				fmt.Print(err)

				db.Query("delete from " + entry2.Text + " where groupMeta_ID=" + strconv.Itoa(ID) + "")

				id.Text = ""
				id.Refresh()

			})

			w.SetContent(
				container.NewVBox(id, submit_btn),
			)
			w.Show()

		}

		//---------------------delete--------------------user_message Table------------------------------------
		if entry2.Text == "user_message" {

			w := a.NewWindow("user_message")
			w.Resize(fyne.NewSize(400, 400))
			id := widget.NewEntry()
			id.SetPlaceHolder("Enter user_message_ID...")

			submit_btn := widget.NewButton("Delete", func() {

				var ID int
				ID, err = strconv.Atoi(id.Text)
				fmt.Print(err)

				db.Query("delete from " + entry2.Text + " where user_message_ID=" + strconv.Itoa(ID) + "")

				id.Text = ""
				id.Refresh()

			})

			w.SetContent(
				container.NewVBox(id, submit_btn),
			)
			w.Show()

		}
		//---------------------delete--------------------user_follower Table------------------------------------
		if entry2.Text == "user_follower" {

			w := a.NewWindow("user_follower")
			w.Resize(fyne.NewSize(400, 400))
			id := widget.NewEntry()
			id.SetPlaceHolder("Enter user_follower_ID...")

			submit_btn := widget.NewButton("Delete", func() {

				var ID int
				ID, err = strconv.Atoi(id.Text)
				fmt.Print(err)

				db.Query("delete from " + entry2.Text + " where user_follower_ID=" + strconv.Itoa(ID) + "")

				id.Text = ""
				id.Refresh()

			})

			w.SetContent(
				container.NewVBox(id, submit_btn),
			)
			w.Show()

		}
	})

	//to clear
	btnClear := widget.NewButton("Clear", func() {
		entry1.SetText("")
		entry2.SetText("")
		result.Text = ""

	})

	w.SetContent(
		container.NewVBox(

			entry1,
			entry2,
			result,
			container.NewGridWithColumns(
				4,
				btn1,
				btn2,
				btn3,
				btn4,
				btn5,
				btnClear,
			),
		),
	)
	w.ShowAndRun()
}
