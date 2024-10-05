//1. GO
//in terms of execution speed - (compilation speed) Rust, C, C++, Java, C# are faster than (execution speed) JS, Py, Ruby, PHP. - Compiled Langages that run on top of a Virtual Machine are - Java, C#, GO.
// GO program uses much less memory but its execution speedmight be a little less fast because of GO runtime.

// package main

// import "fmt"

// func main() { // main() - entry point of program
// 	fmt.Println("hello world!")
// }

// var empty string //is the same as emptyNew := ' '

// newCars := 10 // inferred to be an integer
// temp := 0.0 // temp is inferred to be a floating type data
// var isFuny = true // isFuny is inferred to be a boolean

// // outside of a function (in the global/package scope), every statement begins with a keyword (var, func, and so on) and so the := construct is not availeble.

// package main

// import "fmt"

// func main() {
// 	congrats := 'happy birthday!'
// 	fmt.Println()(congrats)
// }

// package main

// import "fmt"

// func main() {
// 	penniesPerText := 2.0
// 	fmt.Printf('the type of penniesPerText is %T\n', penniesPerText)
// }
// // here, %T is a formatting verb that tells the GO compiler to insert the format/data type of the variable and not its value.

// // Same Line Declarations - we are able to declare multiple variables on the same line -
// milege, company := 38837, 'Tesla'

// // its the same as
// milege := 38837
// company := 'Tesla'

// package main

// import "fmt"

// func main() {
// 	averageOpenRate, displayMessage := 0.23, 'average open rate of your messages'
// 	fmt.Printf(displayMessage, averageOpenRate)
// }

package main

import "fmt"

func main() {
	accountAge := 2.6
	accountAgeInt := int(accountAge)

	fmt.Println("Your acount has existed for", accountAgeInt, "years")
}
