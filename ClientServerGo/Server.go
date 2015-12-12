package main
import (
	"strings"
	"bufio"
	"fmt"
	"net"
	"os"
	"regexp"
	"github.com/mailgun/mailgun-go"
)



func main() {
	ln, _ := net.Listen("tcp", ":8080")
	conn, _ := ln.Accept()
	for true {
		message, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Message Received:", string(message))
		tokens:=strings.Split(message,"%");
		if(tokens[0]=="Add"){
			res:="Yes"
			userinfo:=strings.Split(tokens[1],",")
			if(!verifyEmail(userinfo[2])){
				res="Email not valid or is already taken"
			}else if(!verifyUser(userinfo[0])){
				res="User is already taken"
			}else if(!verifyCedula(userinfo[3])){
				res="Id number not valid or is already taken"
			}else if(!verifyDate(userinfo[4])){
				res="Date not valid or is already taken"
			}
			err:=writeFile(tokens[1])
			if(!err) {
			res = "NO"
			}
			conn.Write([]byte(res+"\n"))
		}
		if(tokens[0]=="Show"){
			res:=search(tokens[1])
			conn.Write([]byte(res+"\n"))
		}
		if(tokens[0]=="Delete"){
			users:=getUsers();
			userlist:= strings.Split(users,";")
			reWriteFile();
			res:="No";
			for i:=0;i<len(userlist);i++{
				tok:=strings.Split(userlist[i],",");
				if(strings.TrimSpace(tok[0])!=strings.TrimSpace(tokens[1])){
					writeFile(userlist[i]);
				}else{
					res="Yes"
				}
			}
			conn.Write([]byte(res+"\n"))

		}
		if(tokens[0]=="Send"){
			//userinfo:= search(tokens[1])
			gun := mailgun.NewMailgun("sandbox001786de44a44eec898cd90610e9097d.mailgun.org", "key-410b4bd9be9ab2241c624fd0a6bd35bf ", "pubkey-ab89ca5403f09e9b41644aba6ee6982b")
			m := mailgun.NewMessage("Sender ClientServe@info.com", "Contact Info", "Message Body", "Recipient "+strings.TrimSpace(tokens[2]))
			response, id, err := gun.Send(m)
			if(err!=nil){
				fmt.Println(err.Error())
			}
			fmt.Printf("Response ID: %s\n", id)
			fmt.Printf("Message from server: %s\n", response)
		}
	}
}

func verifyEmail(email string) bool{
	re:=regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	fmt.Println(re.MatchString(email))
	fmt.Println(isUnique(email,2))
	if(re.MatchString(email)&& isUnique(email,2)){
		return true
	}else{
		return false
	}
}

func verifyCedula(cedula string) bool{
	re:=regexp.MustCompile(`\d{4}-\d{4}-\d{5}`)
	if(re.MatchString(cedula)&& isUnique(cedula,3)){
		return true
	}else{
		return false
	}
}
func verifyDate(date string) bool{
	re:=regexp.MustCompile(`\d{2}-\d{2}-\d{4}`)
	if(re.MatchString(date)){
		return true
	}else{
		return false
	}
}
func verifyUser(user string) bool{
	if(isUnique(user,0)) {
		return true
	}else{
		return false;
	}
}


func writeFile(user string) bool {
	f, err := os.OpenFile("Data.txt", os.O_APPEND ,0666)
	f.WriteString(user+";")
	if(err!=nil) {
		fmt.Println(err.Error())
		return false;
	}
	f.Close();
	return true;
}

func reWriteFile() bool {
	err := os.Remove("Data.txt")
	if(err!=nil) {
		fmt.Println(err.Error())
		return false;
	}
	f,err1:=os.Create("Data.txt");
	if(err1!=nil) {
		fmt.Println(err1.Error())
		return false;
	}
	f.Close();
	return true;
}

func getUsers() string{
	f, err := os.OpenFile("Data.txt", os.O_RDONLY, 0666)
	data:=make([]byte,2042)
	_,err2:=f.Read(data);
	if(err!=nil) {
		fmt.Println(err.Error())
		return " ";
	}
	if(err2!=nil) {
		fmt.Println(err2.Error())
		return " ";
	}
	f.Close();
	return string(data)
}


func search(username string) string{
	users:=getUsers()
	fmt.Println(users)
	userlist:=strings.Split(users,";")
	for i:=0;i<len(userlist);i++{
		userinfo:=strings.Split(strings.TrimSpace(userlist[i]),",")
		if(userinfo[0]==strings.TrimSpace(username)){
			return userlist[i]
		}
	}
	return "No"
}

func isUnique(str string,pos int) bool{
	users:=getUsers()
	userlist:=strings.Split(users,";")
	if(len(userlist)<=1) {
		return true
	}

	for i := 0; i < len(userlist); i++ {
		userinfo:=strings.Split(userlist[i],",")
		if(len(userinfo)<=1) {
			return true;
		}
		if(userinfo[pos]==str) {
			return false;
		}
	}
	return true;
}