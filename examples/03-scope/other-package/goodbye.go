package goodbye

// functions started with capital letter can be used outside of the package
// this is valid for variables too!!
func Say() string {
	return "Goodbye from other package"
}

// since this function is not started with capital letter, it can't be used outside of the package
func scream() string {
	return "GOODBYE FROM OTHER PACKAGE"
}
