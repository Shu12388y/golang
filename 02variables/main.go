package main
import ("fmt")


const LoginToken string = "nsdjnsadjksnakn"  // public variable the first letter should be capital 

func main() {


	// string
	var username string = "shubham"
	fmt.Println(username)
	fmt.Printf("the type of username is %T \n ", username)


	// boolean
	var isLogged bool = true
	fmt.Println(isLogged)
	fmt.Printf("The type of isLogged is %T", isLogged)

	// integer
	var number int = 234
	fmt.Println(number)
	fmt.Printf("The type of number is %T", number)


	// floating
	var floating float32 = 4.34
	fmt.Println(floating)

	
	// declare the variable without var keyword
	jwt := "ndjsnfnnfkf"
	fmt.Println(jwt)
	fmt.Printf("The type of jwt is %T",jwt)

}
