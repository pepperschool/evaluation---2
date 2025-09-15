package main

func split(s string) {
	for _, v := range s {
		print(string(v))
		print("\n")
	}
}

func main() {
	split("LES TEST")
}
