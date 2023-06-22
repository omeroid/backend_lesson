package main

import (
	"bytes"
	"context"
	"fmt"

	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"testing"

	"encoding/json"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/labstack/echo/v4"
	"github.com/omeroid/kosen_backend_lesson/db"
	"github.com/omeroid/kosen_backend_lesson/handler"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type TestSignUpInput struct {
	Username any `json:"userName"`
	Password any `json:"password"`
}
type TestSignInInput struct {
	Username any `json:"userName"`
	Password any `json:"password"`
}

func TestSignUpHandler(t *testing.T) {
	url := "http://127.0.0.1:1323/user/signup"

	e := echo.New()
	e.Logger.SetOutput(ioutil.Discard)
	e.POST("/user/signup", handler.SignUp)

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}

	timeStamp := time.Now()

	mock.ExpectQuery("(?i)SELECT sqlite_version()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.39.4"))
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO `users` (`name`,`password_hash`,`created_at`,`updated_at`,`deleted_at`) VALUES (?,?,?,?,?) RETURNING `id`")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password_hash", "created_at", "updated_at", "deleted_at"}).
			AddRow(2, "wada", "aaf2f198", timeStamp, timeStamp, nil))
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
		number   int
	}{
		"OK": {
			params: TestSignUpInput{
				Username: "wada",
				Password: "aaaa",
			},
			expected: &handler.SignUpOutput{
				ID:        2,
				Name:      "wada",
				CreatedAt: timeStamp,
			},
			number: 1,
		},
		"InvalidUsername": {
			params: TestSignUpInput{
				Username: 1,
				Password: "aaaa",
			},
			expected: &handler.ErrorResponse{
				Message: "code=400, message=Unmarshal type error: expected=string, got=number, field=userName, offset=13, internal=json: cannot unmarshal number into Go struct field SignUpInput.userName of type string (入力値エラー)",
			},
			number: 2,
		},
		"InvalidPassword": {
			params: TestSignUpInput{
				Username: "wada",
				Password: 1,
			},
			expected: &handler.ErrorResponse{
				Message: "code=400, message=Unmarshal type error: expected=string, got=number, field=password, offset=31, internal=json: cannot unmarshal number into Go struct field SignUpInput.password of type string (入力値エラー)",
			},
			number: 3,
		},
		"MissingUsername": {
			params: TestSignUpInput{
				Password: "aaaa",
			},
			expected: &handler.ErrorResponse{
				Message: "(ユーザーネームが空)",
			},
			number: 4,
		},
		"MissingPassword": {
			params: TestSignUpInput{
				Username: "wada",
			},
			expected: &handler.ErrorResponse{
				Message: "(パスワードが空)",
			},
			number: 5,
		},
	}

	for _, v := range vectors {
		fmt.Printf("testcase %d ", v.number)

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

		flag := assert.JSONEq(t, string(expectedJsonData), string(byteArray))
		fmt.Println(flag)
	}

	ctx := context.Background()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}

func TestSignInHandler(t *testing.T) {
	url := "http://127.0.0.1:1323/user/signin"

	e := echo.New()
	e.Logger.SetOutput(ioutil.Discard)
	e.POST("/user/signin", handler.SignIn)

	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Error(err.Error())
	}

	timeStamp := time.Now()
	mock.ExpectQuery("(?i)SELECT sqlite_version()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.39.4"))
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `users` WHERE name=? LIMIT 1")).
		WillReturnRows(sqlmock.NewRows([]string{"id", "name", "password_hash", "created_at", "updated_at", "deleted_at"}).
			AddRow(2, "wada", "aaf2f198", timeStamp, timeStamp, nil))
	mock.ExpectBegin()
	mock.ExpectQuery("DELETE FROM `sessions` WHERE user_id = ?").
		WithArgs(sqlmock.AnyArg).
		WillReturnRows(sqlmock.NewRows([]string{"id", "token", "expire_at", "created_at", "updated_at", "deleted_at"}).
			AddRow(1, "ttttoken", time.Now().Add(24*time.Hour), timeStamp, timeStamp, nil))

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
		params   TestSignInInput
		expected any
		number   int
	}{
		"OK": {
			params: TestSignInInput{
				Username: "wada",
				Password: "aaaa",
			},
			expected: &handler.SignInOutput{},
			number:   1,
		},
		"InvalidUsername": {
			params: TestSignInInput{
				Username: 1,
				Password: "aaaa",
			},
			expected: &handler.ErrorResponse{
				Message: "code=400, message=Unmarshal type error: expected=string, got=number, field=userName, offset=13, internal=json: cannot unmarshal number into Go struct field SignInInput.userName of type string (入力値エラー)",
			},
			number: 2,
		},
		"InvalidPassword": {
			params: TestSignInInput{
				Username: "wada",
				Password: 1,
			},
			expected: &handler.ErrorResponse{
				Message: "code=400, message=Unmarshal type error: expected=string, got=number, field=password, offset=31, internal=json: cannot unmarshal number into Go struct field SignInInput.password of type string (入力値エラー)",
			},
			number: 3,
		},
		"MissingUsername": {
			params: TestSignInInput{
				Password: "aaaa",
			},
			expected: &handler.ErrorResponse{
				Message: "(ユーザーネームが空)",
			},
			number: 4,
		},
		"MissingPassword": {
			params: TestSignInInput{
				Username: "wada",
			},
			expected: &handler.ErrorResponse{
				Message: "(パスワードが空)",
			},
			number: 5,
		},
	}

	for _, v := range vectors {
		fmt.Printf("testcase %d", v.number)

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

		flag := assert.JSONEq(t, string(expectedJsonData), string(byteArray))
		fmt.Println(flag)
	}

}
