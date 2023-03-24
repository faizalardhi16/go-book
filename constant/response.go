package constant

type ResponseType struct {
	NotFound string
	Failed   string
	Success  string
}

var ResponseEnum = ResponseType{
	NotFound: "Not Found",
	Failed:   "Failed",
	Success:  "Success",
}
