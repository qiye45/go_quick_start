package main

func main() {
	defer func() {
		defer recover()
	}()

	panic("it is panic") // recover
}
