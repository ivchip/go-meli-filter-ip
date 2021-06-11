package http_test

import (
	"github.com/ivchip/go-meli-filter-ip/domain/fixtures"
	"github.com/ivchip/go-meli-filter-ip/domain/mocks"
	delivery "github.com/ivchip/go-meli-filter-ip/interface/http"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Handlers", func() {
	Describe("ContextualResultHandlers", func() {
		var mockAccountUseCases mocks.MockContextualResultUseCases
		BeforeEach(func() {
			mockAccountUseCases = mocks.MockContextualResultUseCases{}
		})

		Context("Given a valid IP", func() {
			It("should give code 200", func() {
				ip := fixtures.GenerateIP()
				// Arrange
				mockAccountUseCases.On("GetByIP", ip).
					Return(fixtures.GenerateContextualResult(), nil)
				req := httptest.NewRequest("GET", "/getIP", nil)
				resp := httptest.NewRecorder()

				// Act
				delivery.NewMuxChiRouter().ServeHTTP(resp, req)

				// Asserts
				Expect(resp.Code).To(Equal(http.StatusOK))
			})
		})
	})

})
