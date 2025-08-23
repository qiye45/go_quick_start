package main

func main() {
	defer func() {
		recover()
	}()

	panic("it is panic") // recover
}
