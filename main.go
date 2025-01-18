package main

type HTTPResponse interface {
	GetStatusCode() int
	GetBody() struct{}
}

type HTTPClient interface {
	Post(url string, body struct{}) HTTPResponse
	Get(url string) HTTPResponse
}

func main() {

}
