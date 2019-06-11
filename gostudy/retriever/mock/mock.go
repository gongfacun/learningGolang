package mock

import "fmt"

type Retriever struct {
	Conttexts string
}

func (r *Retriever) Get(url string) string {
	return r.Conttexts
}

//Post
func (r *Retriever) Post(url string, form map[string]string) string {
	r.Conttexts = form["abc"]
	return "ok"
}

func (r *Retriever) String() string {
	return fmt.Sprintf("mockRetriever:%s", r.Conttexts)
}
