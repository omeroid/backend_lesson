package main

import (
	"bytes"
	"fmt"

	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/omeroid/kosen_backend_lesson/db"
	"github.com/omeroid/kosen_backend_lesson/handler"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"k8s.io/apimachinery/pkg/util/json"
)

type TestSignUpInput struct {
	Username any `json:"userName"`
	Password any `json:"password"`
}

func TestSignUpHandler(t *testing.T) {
	e := echo.New()
	e.Logger.SetOutput(ioutil.Discard)
	e.POST("/user/signup", handler.SignUp)

	url := "http://127.0.0.1:1323/user/signup"

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}

	createdAt := time.Now()
	updatedAt := time.Now()
	mock.ExpectQuery("(?i)SELECT sqlite_version()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.39.4"))
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password_hash`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?) RETURNING `id`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password_hash", "created_at", "updated_at", "deleted_at"}).
			AddRow(2, "wada", "aaf2f198", createdAt, updatedAt, nil))
	mock.ExpectCommit()
	mockDB, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite3",
		Conn:       sqlDB,
	}, &gorm.Config{})
	if err != nil {
		t.Error(err.Error())
	}

	e.Use(db.DBMiddleware(mockDB))

	go e.Start(":1323")

	vectors := map[string]struct {
		params   TestSignUpInput
		expected any
	}{
		"OK": {
			params: TestSignUpInput{
				Username: "wada",
				Password: "aaaa",
			},
			expected: &handler.SignUpOutput{
				ID:        2,
				Name:      "wada",
				CreatedAt: createdAt,
			},
		},
		"InvalidUsername": {
			params: TestSignUpInput{
				Username: 1,
				Password: "aaaa",
			},
			expected: &handler.ErrorResponse{
				Message: "username invalid",
			},
		},
		"InvalidPassword": {
			params: TestSignUpInput{
				Username: "wada",
				Password: 1,
			},
			expected: &handler.ErrorResponse{
				Message: "password invalid",
			},
		},
	}

	for _, v := range vectors {
		jsonData, err := json.Marshal(v.params)
		if err != nil {
			t.Error(err.Error())
		}
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
		if err != nil {
			t.Error(err)
		}
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Error(err)
		}
		defer resp.Body.Close()

		expectedJsonData, err := json.Marshal(v.expected)
		if err != nil {
			t.Error(err)
		}

		byteArray, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Error(err)
		}

		flag := assert.JSONEq(t, string(byteArray), string(expectedJsonData))
		fmt.Println(flag)
	}
}
