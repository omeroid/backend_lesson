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
	//新規のSignUpInputオブジェクトを用意した上で、リクエスト内のデータをそのオブジェクトにバインドします。
	input := new(SignUpInput)
	if err := c.Bind(input); err != nil {
		//エラーが発生した場合、400エラーと共にエラーメッセージを返却します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (入力値エラー)", err),
		})
	}
	//DBのコネクションを取得
	// echo.ContextからDBへのコネクションを取得します。
	// キャストしており、取得したインターフェースを正しい型 (*gorm.DB) に変換しています。
	conn := c.Get("db").(*gorm.DB)
	//パスワードのHash化処理
	//入力されたパスワードをbcryptを使ってハッシュ化します。ハッシュ化際のコストは10です。
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), 10)
	if err != nil {
		//エラーが発生した場合、400エラーと共にエラーメッセージを返却します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (hash化したpasswordの生成に失敗)", err),
		})
	}
	//usersテーブルにレコードを追加
	//Userの新インスタンスを作成し、ユーザ名とハッシュ化したパスワードを設定して、データベースにユーザ情報を保存します。
	user := db.User{
		Name:         input.Username,
		PasswordHash: string(hashedPassword),
	}
	if result := conn.Create(&user); result.Error != nil {
		//エラーが発生した場合、400エラーと共にエラーメッセージを返却します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (user作成エラー)", result.Error),
		})
	}
	//JSONレスポンスに返却するデータを設定します。新規作成されたユーザのID、名前、登録日時を含めます。
	output := SignUpOutput{
		ID:        user.ID,
		Name:      user.Name,
		CreatedAt: user.CreatedAt,
	}
	//ステータスコード201と出力内容を返却します。
	return c.JSON(http.StatusCreated, output)
}

// ユーザを認証する
func SignIn(c echo.Context) error {
	//入力値の取得
	//新規のSignInInputオブジェクトを用意した上で、リクエスト内のデータをそのオブジェクトにバインドします。
	input := new(SignInInput)
	if err := c.Bind(input); err != nil {
		//エラーが発生した場合、400エラーと共にエラーメッセージを返却します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (入力値エラー)", err),
		})
	}
	//DBのコネクションを取得
	// echo.ContextからDBへのコネクションを取得します。
	// キャストしており、取得したインターフェースを正しい型 (*gorm.DB) に変換しています。
	conn := c.Get("db").(*gorm.DB)
	//usersテーブルにusernameで検索をかける
	//入力されたユーザ名でusersテーブルを検索します。
	user := db.User{}
	if result := conn.Take(&user, "name=?", input.Username); result.Error != nil {
		//エラーが発生した場合、400エラーと共にエラーメッセージを返却します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (user検索エラー)", result.Error),
		})
	}
	//過去のsessionの削除
	//ユーザIDに紐づく過去のセッションを削除します。
	if result := conn.Where("user_id = ?", user.ID).Delete(&db.Session{}); result.Error != nil {
		//エラーが発生した場合、400エラーと共にエラーメッセージを返却します。
		return c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (過去session削除エラー)", result.Error),
		})
	}
	//パスワードの比較(Hash化されている)
	//入力されたパスワードがDB内のパスワードハッシュと一致するかチェックします。
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(input.Password)); err != nil {
		//エラーが発生した場合、401エラーと共にエラーメッセージを返却します。
		return c.JSON(http.StatusUnauthorized, ErrorResponse{
			Message: fmt.Sprintf("%s (passwordが違う)", err),
		})
	}
	//sessionsテーブルに登録するtokenの生成（uuid）
	//新規のUUIDを生成します。これがセッショントークンとして使用されます。
	token, err := uuid.NewRandom()
	if err != nil {
		//エラーが発生した場合、500エラーと共にエラーメッセージを返却します。
		return c.JSON(http.StatusInternalServerError, ErrorResponse{
			Message: fmt.Sprintf("%s (token生成エラー)", err),
		})
	}
	//Sessionへの追加
	//新たなセッション情報をデータベースに保存します。ユーザID、トークン、有効期限を設定します。
	session := db.Session{
		UserID:    user.ID,
		Token:     token.String(),
		ExpiredAt: time.Now().Add(time.Hour * 24),
	}
	if result := conn.Create(&session); result.Error != nil {
		//エラーが認証した場合、400エラーと共にエラーメッセージを返却します。
		c.JSON(http.StatusBadRequest, ErrorResponse{
			Message: fmt.Sprintf("%s (session作成エラー)", result.Error),
		})
	}
	//JSONレスポンスに返却するデータを設定します。認証されたユーザID、ユーザ名、セッショントークンを設定します。
	output := SignInOutput{
		UserID:   user.ID,
		UserName: user.Name,
		Token:    token.String(),
	}
	//ステータスコード200と出力内容を返却します。
	return c.JSON(http.StatusOK, output)
}

// userにsessionが存在するか確認した後失効していないか確認する
func IsSessionValid(conn *gorm.DB, token string) error {
	//sessionsをtokenで検索する
	//入力されたトークンでsessionsテーブルを検索します。
	session := db.Session{}
	if result := conn.First(&session, "token = ?", token); result.Error != nil {
		//エラーが発生した場合、エラーメッセージを返却します。
		return errors.New(result.Error.Error() + " (sessionの検索エラー)")
	}
	//tokenの有効期限が切れている時
	//セッションの有効期限が現在時刻より以前になっている場合、エラーメッセージを返却します。
	if session.ExpiredAt.Before(time.Now()) {
		return errors.New("session expired" + " ログインし直してください")
	}
	return nil
}
