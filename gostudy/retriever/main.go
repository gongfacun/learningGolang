package main

import (
	"fmt"
	"gostudy/retriever/mock"
	"gostudy/retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}
type Poster interface {
	Post(url string, form map[string]string) string
}
type RetrieverPoster interface {
	Retriever
	Poster
}

func session(rp RetrieverPoster) string {
	rp.Post("hhahahah", map[string]string{

		"dafd": "ddfsdfs",
		"abc":  "www.haowan.com",
	})
	return rp.Get("new url hahah")
}

func dowload(r Retriever) string {
	return r.Get("http://www.mooc.com")
}

func post(p Poster) string {

	return p.Post("http://www.mooc.com", map[string]string{
		"firstname": "wenjianwenjian",
		"course":    "lalalala",
		"abc":       "www.hahahh.com",
	})
}

func main() {
	var tm Retriever
	hah := &mock.Retriever{"sfsdfasfasf"}
	tm = hah
	fmt.Printf("%T   %V \n", tm, tm)
	fmt.Println(session(hah))
	fmt.Println(post(hah))

	tm = &real.Retriever{

		UserAgent: "Mallize/5.0",
		TimeOut:   time.Minute,
	}
	fmt.Printf("%T  %V \n", tm, tm)
	switch v := tm.(type) {
	case *mock.Retriever:
		fmt.Println("mock.retriever:", v.Conttexts)
	case *real.Retriever:
		fmt.Println("real zhichen:", v.UserAgent)

	}
	//fmt.Println(dowload(tm))
	fmt.Println(hah)
}
