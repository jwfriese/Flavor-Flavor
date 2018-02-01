package flayfine_acceptance_tests

import (
	"bytes"
	"io/ioutil"
	"net/http"

	. "github.com/Benjamintf1/Expanded-Unmarshalled-Matchers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("/flavors", func() {
	var (
		url string
	)

	BeforeEach(func() {
		url = "http://localhost:8181/flavors"
	})

	Describe("Getting flavors", func() {
		var (
			initialFlavorsResponseBody []byte
		)

		BeforeEach(func() {
			response, err := http.Get(url)

			Expect(err).ToNot(HaveOccurred())

			initialFlavorsResponseBody, err = ioutil.ReadAll(response.Body)
			response.Body.Close()

			Expect(err).ToNot(HaveOccurred())
		})

		It("returns initial list of flavors", func() {
			Expect(initialFlavorsResponseBody).To(MatchUnorderedJSON([]byte(`
				[
					{ "name": "garlic" },
					{ "name": "tomato" },
					{ "name": "thyme" },
					{ "name": "vanilla" },
					{ "name": "lemon" },
					{ "name": "rosemary" }
				]
		`)))
		})

		Describe("Creating new flavors", func() {
			var (
				followUpFlavorResponseBody []byte
			)

			BeforeEach(func() {
				createFlavorRequestBody := bytes.NewBuffer([]byte(`
				{
					"name": "almond"
				}`))

				request, createRequestErr := http.NewRequest("POST", url, createFlavorRequestBody)
				request.Header.Add("Content-Type", "application/json")

				Expect(createRequestErr).ToNot(HaveOccurred())

				response, err := http.DefaultClient.Do(request)

				Expect(err).ToNot(HaveOccurred())

				response, err = http.Get(url)

				Expect(err).ToNot(HaveOccurred())

				followUpFlavorResponseBody, err = ioutil.ReadAll(response.Body)
				response.Body.Close()

				Expect(err).ToNot(HaveOccurred())
			})

			It("returns updated list of flavors", func() {
				Expect(followUpFlavorResponseBody).To(MatchUnorderedJSON([]byte(`
				[
					{ "name": "garlic" },
					{ "name": "tomato" },
					{ "name": "thyme" },
					{ "name": "vanilla" },
					{ "name": "lemon" },
					{ "name": "almond" },
					{ "name": "rosemary" }
				]`)))
			})
		})
	})
})
