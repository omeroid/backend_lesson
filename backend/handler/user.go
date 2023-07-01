package handler

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/omeroid/backend_lesson/backend/pkg/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// userを登録する
func SignUp(c echo.Context) error {
	//入力値の取得
	input := new(SignUpInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (入力値エラー)", err),
		})
	}

	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//パスワードのHash化処理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (hash化したpasswordの生成に失敗)", err),
		})
	}

	//usersテーブルにレコードを追加
	user := db.User{
		Name:         input.Username,
		PasswordHash: string(hashedPassword),
	}
	if result := conn.Create(&user); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (user作成エラー)", result.Error),
		})
	}

	output := SignUpOutput{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}

	return c.JSON(http.StatusCreated, output)
}

// ユーザを認証する
func SignIn(c echo.Context) error {
	//入力値の取得
	input := new(SignInInput)
	if err := c.Bind(input); err != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (入力値エラー)", err),
		})
	}

	//DBのコネクションを取得
	conn := c.Get("db").(*gorm.DB)

	//usersテーブルにusernameで検索をかける
	user := db.User{}
	if result := conn.Take(&user, "name=?", input.Username); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (user検索エラー)", result.Error),
		})
	}

	//過去のsessionの削除
	if result := conn.Where("user_id = ?", user.ID).Delete(&db.Session{}); result.Error != nil {
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (過去session削除エラー)", result.Error),
		})
	}

	//パスワードの比較(Hash化されている)
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (passwordが違う)", err),
		})
	}

	//sessionsテーブルに登録するtokenの生成（uuid）
	token, err := uuid.NewRandom()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: fmt.Sprintf("%s (token生成エラー)", err),
		})
	}

	//Sessionへの追加
	session := db.Session{
		UserID:    user.ID,
		Token:     token.String(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}
	if result := conn.Create(&session); result.Error != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (session作成エラー)", result.Error),
		})
	}

	output := SignInOutput{
		UserID:   user.ID,
		UserName: user.Name,
		Token:    token.String(),
	}

	return c.JSON(http.StatusOK, output)
}

// userにsessionが存在するか確認した後失効していないか確認する
func IsSessionValid(conn *gorm.DB, token string) error {
	//sessionsをtokenで検索する
	session := db.Session{}
	if result := conn.First(&session, "token = ?", token); result.Error != nil {
		return errors.New(result.Error.Error() + " (sessionの検索エラー)")
	}

	//tokenの有効期限が切れている時
	if session.ExpiredAt.Before(time.Now()) {
		return errors.New("session expired" + " ログインし直してください")
	}

	return nil
}
