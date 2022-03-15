package server

import (
	"log"
	"net/http"

	"22dojo-online/pkg/adapter/controllers"
	"22dojo-online/pkg/adapter/middleware"
	"22dojo-online/pkg/domain/service"
	driver "22dojo-online/pkg/driver/mysql"
	"22dojo-online/pkg/driver/mysql/database"
	"22dojo-online/pkg/usecase"
)

// Serve HTTPサーバを起動する
func Serve(addr string) {
	db := driver.NewSQLHandler()
	userRepo := database.NewUserRepository(db)
	userService := service.NewUserService(userRepo)
	userInteractor := usecase.NewUserInteractor(userService)
	userController := controllers.NewUserController(userInteractor)
	m := middleware.NewAuth(userInteractor)

	/* ===== URLマッピングを行う ===== */
	http.HandleFunc("/setting/get", get(controllers.GetSetting()))

	http.HandleFunc("/user/get",
		get(m.Authenticate(userController.GetUser())))
	http.HandleFunc("/user/create",
		post(userController.InsertUser()))
	http.HandleFunc("/user/update",
		post(m.Authenticate(userController.UpdateUser())))

	/* ===== サーバの起動 ===== */
	log.Println("Server running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("Listen and serve failed. %+v", err)
	}
}

// get GETリクエストを処理する
func get(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodGet)
}

// post POSTリクエストを処理する
func post(apiFunc http.HandlerFunc) http.HandlerFunc {
	return httpMethod(apiFunc, http.MethodPost)
}

// httpMethod 指定したHTTPメソッドでAPIの処理を実行する
func httpMethod(apiFunc http.HandlerFunc, method string) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		// CORS対応
		writer.Header().Add("Access-Control-Allow-Origin", "*")
		writer.Header().Add("Access-Control-Allow-Headers", "Content-Type,Accept,Origin,x-token")

		// プリフライトリクエストは処理を通さない
		if request.Method == http.MethodOptions {
			return
		}
		// 指定のHTTPメソッドでない場合はエラー
		if request.Method != method {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			if _, err := writer.Write([]byte("Method Not Allowed")); err != nil {
				log.Println(err)
			}
			return
		}

		// 共通のレスポンスヘッダを設定
		writer.Header().Add("Content-Type", "application/json")
		apiFunc(writer, request)
	}
}
