package main_test

import (
	"encoding/json"
	. "github.com/nonfu/gwp/webservice/rest"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"net/http"
	"net/http/httptest"
)

var _ = Describe("Rest", func() {
	var mux *http.ServeMux
	var post *FakePost
	var writer *httptest.ResponseRecorder

	BeforeEach(func() {
		post = &FakePost{}
		mux = http.NewServeMux()
		mux.HandleFunc("/post", HandleRequest(post))
		writer = httptest.NewRecorder()
	})

	Context("Get a post using id", func() {
		It("Should get a post", func() {
			request, _ := http.NewRequest("GET", "/post?id=1", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(200))

			var post Post
			json.Unmarshal(writer.Body.Bytes(), &post)
			Expect(post.Id).To(Equal(1))
		})
	})

	Context("Get an error if post id is not an integer", func() {
		It("should get a HTTP 500 response", func() {
			request, _ := http.NewRequest("GET", "/post?id=hello", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(500))
		})
	})
})

type FakePost struct {
	Id      int
	Title   string
	Content string
	Author  string
}

func (post *FakePost) Retrieve(id int) (err error) {
	post.Id = id
	return
}

func (post *FakePost) Create() (err error) {
	return
}

func (post *FakePost) Update() (err error) {
	return
}

func (post *FakePost) Delete() (err error) {
	return
}
