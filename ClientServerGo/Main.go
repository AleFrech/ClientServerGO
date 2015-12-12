package main
import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"net"
)

func main() {
	conn, _ := net.Dial("tcp", "127.0.0.1:8080")
	for true {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("1) Add User")
		fmt.Println("2) Show User")
		fmt.Println("3) Delete User")
		fmt.Println("4) Send User")
		option,_:= reader.ReadString('\n');
		switch(strings.TrimSpace(option)){
		case"1":
			fmt.Println("Enter Username")
			userename,_:= reader.ReadString('\n')
			userename=strings.TrimSpace(userename)
			fmt.Println("Enter Name")
			name,_:= reader.ReadString('\n')
			name=strings.TrimSpace(name)
			fmt.Println("Enter Email")
			email,_:= reader.ReadString('\n')
			email=strings.TrimSpace(email)
			fmt.Println("Enter Id Number")
			cedula,_:= reader.ReadString('\n')
			cedula=strings.TrimSpace(cedula)
			fmt.Println("Enter Date")
			date,_:= reader.ReadString('\n')
			date=strings.TrimSpace(date)
			fmt.Println("Enter img url")
			img,_:= reader.ReadString('\n')
			img=strings.TrimSpace(img);
			user:=userToString(userename,name,email,cedula,date,img)
			userToAdd:="Add"+"%"+user+"\n"
			conn.Write([]byte(userToAdd))
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Println("Message Received:", string(message))
			if(strings.TrimSpace(string(message))=="Yes"){
				fmt.Println("User succesfully Added!!!!")
			}else{
				fmt.Println("Error User not Added!!!!!!!")
			}
			break;
		case "2":
			fmt.Println("Enter Username")
			userename,_:= reader.ReadString('\n')
			userename=strings.TrimSpace(userename)
			userToShow:="Show"+"%"+userename+"\n"
			conn.Write([]byte(userToShow));
			message, _ := bufio.NewReader(conn).ReadString('\n')
			if(strings.TrimSpace(string(message))!="No"){
				userinfo:=strings.Split(message,",")
				println("----Contact Info------ ")
				println("Username : "+userinfo[0])
				println("Name : "+userinfo[1])
				println("Email : "+userinfo[2])
				println("Identity Card : "+userinfo[3])
				println("Birth Date  : "+userinfo[4])
			}else{
				fmt.Println("Error User not in database!!!!!")
			}
			break;
		case "3":
			fmt.Println("Enter Username")
			userename,_:= reader.ReadString('\n')
			userename=strings.TrimSpace(userename)
			userToDelete:="Delete"+"%"+userename+"\n"
			conn.Write([]byte(userToDelete))
			message, _ := bufio.NewReader(conn).ReadString('\n')
			if(strings.TrimSpace(string(message))=="Yes") {
				fmt.Println("User Succesfully Deleted")
			}else{
				fmt.Println("Error User not deleted!!!!!")
			}
			break;
		case "4":
			fmt.Println("Enter Username to send")
			userename,_:= reader.ReadString('\n')
			userename=strings.TrimSpace(userename)
			fmt.Println("Enter Email to send")
			email,_:= reader.ReadString('\n')
			email=strings.TrimSpace(email)
			userToSend:="Send"+"%"+userename+"%"+email+"\n"
			conn.Write([]byte(userToSend))
			message, _ := bufio.NewReader(conn).ReadString('\n')
			if(strings.TrimSpace(string(message))=="Yes") {
				fmt.Println("User Succesfully Sent!!!!!")
			}else{
				fmt.Println("Email not sent !!!!!")
			}
			break;
		case "5":
			os.Exit(1);
			break;
		}
	}
}

func userToString(usr string,name string,email string,cedula string,date string,imgurl string) string{
	return usr+","+name+","+email+","+cedula+","+date+","+imgurl;
}

