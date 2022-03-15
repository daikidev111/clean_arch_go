package test

// func TestSelectUserByPrimaryKeyService(t *testing.T) {
// 	table := []struct {
// 		testName   string
// 		id         string
// 		user       entity.User
// 		wantStatus int
// 	}{
// 		{
// 			"FIRST TEST CASE: TestSelectUserByPrimaryKeyService",
// 			"78164dcf-6b7c-45e4-862a-2a0f6735a449",
// 			entity.User{
// 				ID:        "78164dcf-6b7c-45e4-862a-2a0f6735a449",
// 				AuthToken: "b187b9e0-08e6-42dd-a9b3-a900b137983c",
// 				Name:      "whatt",
// 				HighScore: 100,
// 				Coin:      10000,
// 			},
// 			200,
// 		},
// 		{
// 			"SECOND TEST CASE: TestSelectUserByPrimaryKeyService",
// 			"829bdb53-f322-40bb-9327-63ab00536cd3",
// 			entity.User{
// 				ID:        "829bdb53-f322-40bb-9327-63ab00536cd3",
// 				AuthToken: "4ed9e2e7-cbba-4ab9-a669-14b688dfb245",
// 				Name:      "bruh",
// 				HighScore: 0,
// 				Coin:      0,
// 			},
// 			200,
// 		},
// 		{
// 			"THIRD TEST CASE: TestSelectUserByPrimaryKeyService",
// 			"909b42f8-cce9-4d02-bce4-e7c7e28df550",
// 			entity.User{
// 				ID:        "909b42f8-cce9-4d02-bce4-e7c7e28df550",
// 				AuthToken: "bb68df68-964e-4f27-a225-0cbafdd6ce9f",
// 				Name:      "9",
// 				HighScore: 70,
// 				Coin:      0,
// 			},
// 			200,
// 		},
// 	}
// 	for _, tt := range table {
// 		t.Run(tt.testName, func(t *testing.T) {
// 			r := httptest.NewRequest(http.MethodGet, "/user/get", nil)

// 			query := r.URL.Query()
// 			query.Add("x-token", tt.user.AuthToken)
// 			// got := httptest.NewDecoder()
// 			// con.DetailBlogView(got, r)
// 			// assert.Equal(t, tt.wantStatus, got.Result().StatusCode)
// 			// ctx := request.Context()
// 			// userID := dcontext.GetUserIDFromContext(ctx)
// 		})
// 	}
// }
