/* const => If a variable should have a fixed value that cannot 
			be changed, you can use the const keyword.
			Constant names are usually written in uppercase letters.

			const CONSTNAME type = value
*/

package main
import ("fmt")

const PI = 3.14
const A int = 1

// multiple constant declaration
const (
	D int = 1
	B int = 2
	C int = 3
)

func main() {
	fmt.Println(PI);
	// A = 2 cannot be reassigned
	fmt.Println(A);
	fmt.Println(B);
	fmt.Println(C);
	fmt.Println(D);
	
}