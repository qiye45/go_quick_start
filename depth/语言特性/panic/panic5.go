package main

func main() {
	defer func() {
		defer func() {
			recover()
		}()
	}()

	panic("it is panic") // not recover
}
