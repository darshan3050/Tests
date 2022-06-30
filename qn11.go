
func main() {
	s := make([]int, 0, 5)
	s = append(s, 1, 2, 3, 4)
	a := append(s, 5)
	fmt.Println(a) // O/P :- [1,2,3,4,5]____________5 is appended a________
	b := append(s, 6)
	fmt.Println(b) // O/P :- [1,2,3,4,6]___________6 is apended in b_________
	fmt.Println(a) // O/P :- _[1,2,3,4,5]__________last apend value in a was 5_________
	fmt.Println(s) // O/P :- ____[1,2,3,4]________s has original values________

	// What is the expected output and why?

}
