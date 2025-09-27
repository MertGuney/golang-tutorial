package main

import "fmt"

func main() {
	var encoder IEncoder
	
	encoder = &XEncoder{}
	encoder.Encode("abc")
	encoder.Decode("abc")

	encoder = &YEncoder{}
	encoder.Encode("xyz")
	encoder.Decode("xyz")
}

type IEncoder interface {
	Encode(data string)
	Decode(data string)
}

type XEncoder struct {
}

type YEncoder struct {
}

func (xEncoder *XEncoder) Encode(data string) {
	fmt.Println("Encoding data using XEncoder:", data)
}

func (xEncoder *XEncoder) Decode(data string) {
	fmt.Println("Decoding data using XEncoder:", data)
}

func (yEncoder *YEncoder) Encode(data string) {
	fmt.Println("Encoding data using YEncoder:", data)
}

func (yEncoder *YEncoder) Decode(data string) {
	fmt.Println("Decoding data using YEncoder:", data)
}