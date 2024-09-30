/* Go has three output functions
			
			Print() => prints its arguments with their default format
			Println() => whitespace is added bw the arguments & newline is added at end
			Printf() => formats its arguments based on given formatting verb

			%v	Prints the value in the default format
			%#v	Prints the value in Go-syntax format
			%T	Prints the type of the value
			%%	Prints the % sign

*/

package main
import ("fmt")

func main() {

	//Print()
	var i, j string = "Hello", "World";
	var a, b int = 10, 20;
	fmt.Print(i, " ", j);
	fmt.Print(i, "\n");
	fmt.Print(j, "\n"); 
	fmt.Print(a, b);

	//Println()
	fmt.Println(i, j);

	//Printf()
	fmt.Printf("i has value: %v and type: %T\n", i, i);
	fmt.Printf("a has value: %v and type: %T\n", a, a);
}
