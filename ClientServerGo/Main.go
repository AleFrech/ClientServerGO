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
		if (strings.TrimSpace(option)=="1"){
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
		}
		if (strings.TrimSpace(option)=="2"){
			fmt.Println("Enter Username")
			userename,_:= reader.ReadString('\n')
			userename=strings.TrimSpace(userename)
			userToAdd:="Show"+"%"+userename+"\n"
			conn.Write([]byte(userToAdd))
			message, _ := bufio.NewReader(conn).ReadString('\n')
			fmt.Println("Message Received:", string(message))
			if(strings.TrimSpace(string(message))!=" "){
				userinfo:=strings.Split(message,",")

			}else{
				fmt.Println("Error User not in database!!!!!")
			}
		}
		if (strings.TrimSpace(option)=="3"){

		}
		if (strings.TrimSpace(option)=="4"){

		}
		if (strings.TrimSpace(option)=="5"){
			os.Exit(1);
		}
	}
}

func userToString(usr string,name string,email string,cedula string,date string,imgurl string) string{
	return usr+","+name+","+email+","+cedula+","+date+","+imgurl;
}

