package jsonparser

func Parse(json_string string, ch chan Token) {
	ch <- Tokenize(json_string)
}
