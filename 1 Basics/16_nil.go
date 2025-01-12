package main

// In other programming languages, uninitialized objects are default to null or nil
// But, it doesn't work that way in go
// The default value of a string is "", an int is 0, a boolean is false
// But for pointers, interfaces, functions, maps, slices, and channels, it is nil and it can only be used for those types

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	}

	return map[string]string{
		"name": name,
	}
}

func main() {

}
