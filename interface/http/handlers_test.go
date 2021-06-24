package http_test

import (
	"github.com/ivchip/go-meli-filter-ip/domain/fixtures"
	"github.com/ivchip/go-meli-filter-ip/domain/mocks"
	delivery "github.com/ivchip/go-meli-filter-ip/interface/http"
	"net/http"
	"net/http/httptest"
	"testing"

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

func TestGetByIp(t *testing.T) {
	//ip := "192.146.146.164"
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ipBlocking/:ip", nil)
	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(delivery.GetByIp)
	handler.ServeHTTP(resp, req)
}

func TestGetAll(t *testing.T) {
	resp := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodGet, "/ipBlocking", nil)
	req.Header.Set("Content-Type", "application/json")
	handler := http.HandlerFunc(delivery.GetAll)
	handler.ServeHTTP(resp, req)
}

func TestCreate(t *testing.T) {

}

func TestUpdate(t *testing.T) {

}
