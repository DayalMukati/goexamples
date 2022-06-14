package main
import "fmt"


func main() {
	var x int = 30
	var ip *int
	ip = &x


	fmt.Printf("Address of x variable: ", &x)


	fmt.Printf("Addrsss stored in IP:", ip)

	fmt.Printf("Value of *ip :", *ip)
}