package main

// func TestRecordingWinsAndRetrievingThem(t *testing.T) {
// 	store := NewBoltScoreStore()
// 	server := GoBoardServer{store}
// 	player := "Pepper"

// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))
// 	server.ServeHTTP(httptest.NewRecorder(), newPostWinRequest(player))

// 	response := httptest.NewRecorder()
// 	server.ServeHTTP(response, newGetScoreRequest(player))
// 	assertStatus(t, response.Code, http.StatusOK)

// 	assertResponseBody(t, response.Body.String(), "3")
// }
