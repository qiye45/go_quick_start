package main

func byteArrayToString(b []byte) string {
	return string(b)
}

func stringToByteArray(s string) []byte {
	return []byte(s)
}

func main() {
}

//go tool compile -N -l -S main.go
