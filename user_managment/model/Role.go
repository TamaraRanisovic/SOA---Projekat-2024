package model

// Role - Custom type to hold value for each role 1-3
type Role int

// Declare related constants for each role starting with index 1
const (
	Administrator Role = iota + 1 // EnumIndex = 1
	Guide                       // EnumIndex = 2
	Tourist                      // EnumIndex = 3
)

// String - Creating common behavior - give the type a String function
func (r Role) String() string {
	return [...]string{"Administrator", "Guide", "Tourist"}[r-1]
}

// EnumIndex - Creating common behavior - give the type a EnumIndex functio
func (r Role) EnumIndex() int {
	return int(r)
}

//USAGE EXAMPLE
//func main() {
//	var r Role = Administrator
//	fmt.Println(r)             // Print : Administrator
//	fmt.Println(r.String())    // Print : Administrator
//	fmt.Println(d.EnumIndex()) // Print : 1
//}