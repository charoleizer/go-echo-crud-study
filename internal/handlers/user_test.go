package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/charoleizer/go-echo-crud-study/internal/cache"
	"github.com/charoleizer/go-echo-crud-study/internal/models"

	"github.com/labstack/echo/v4"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

const (
	mockIDJon   = 1
	mockNameJon = "Jon Smith"
)

var (
	mockObject = &models.User{mockIDJon, mockNameJon}
	userJSON   = fmt.Sprintf(`{"id":1,"name":"%s"}%s`, mockNameJon, "\n")
)

func TestUserr(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "User Suite")
}

var _ = When("creating a new user", func() {
	CreateUserSetup := func() (*httptest.ResponseRecorder, echo.Context, *user) {
		e := echo.New()
		request := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(userJSON))
		request.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		respRec := httptest.NewRecorder()
		ctx := e.NewContext(request, respRec)
		hdlr := &user{mockObject}

		return respRec, ctx, hdlr
	}

	Context("given an invalid json object' ", func() {
		_, ctx, hdlr := CreateUserSetup()

		It("should return an error", func() {
			Expect(hdlr.CreateUser(ctx, nil)).To(MatchError("code=400, message=json: Unmarshal(nil *models.User), internal=json: Unmarshal(nil *models.User)"))
		})
	})

	Context("given the name of 'Jon Smith' ", func() {
		respRec, ctx, hdlr := CreateUserSetup()

		It("should create an user without error", func() {
			Expect(hdlr.CreateUser(ctx, &models.User{ID: cache.Seq})).To(BeNil())
		})

		It("should response with a StatusCreated code (201)", func() {
			Expect(respRec.Code).To(Equal(http.StatusCreated))
		})

		It("should response a body equals to mocked userJSON", func() {
			Expect(respRec.Body.String()).To(Equal(userJSON))
		})
	})

})
