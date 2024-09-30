/* Variables => int
				float32
				string
				bool

				var variablename type = value
				var variablename = value
				varialbename := value
*/

package main;
import("fmt");

var outside1 int;
var outside2 int = 2;
var outside3 = 3;

//outsideError := 1; // := cannot be used outside a function and be accessed

func main() {

	//with the var keyword
	var student1 string = "Jesse";
	var student2 string = "Walter";

	//with the := sign
	x := 2; // type is inferred

	fmt.Println(student1);
	fmt.Println(student2);
	fmt.Println(x);

	//default values
	var a string;
	var b int;
	var c bool;

	fmt.Println(a); // ""
	fmt.Println(b); // 0
	fmt.Println(c); // false

	//value assignment after declaration
	var car string;
	car = "Aston Martin"
	fmt.Println(car);

	//accessing variables outside of a function
	outside1 = 1;
	fmt.Println(outside1);
	fmt.Println(outside2);
	fmt.Println(outside3);

	//fmt.Println(outsideError);

	//multiple variable declaration
	var m, n, o, p int = 1, 3, 5, 7;
	fmt.Println(m);
	fmt.Println(n);
	fmt.Println(o);
	fmt.Println(p);

	//multiple variable declaration without type specified
	var r, s = 6, "Hello"
	t, u := 7, "World"

	//variable declaration in a block
	var(
		a int
		b int = 1
		c string = "hello"
	)

}

