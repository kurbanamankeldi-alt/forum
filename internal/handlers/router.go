package handlers

import "net/http"

func (app *App) Router(mux *http.ServeMux) {
		//register endpoints here
	mux.HandleFunc("/profile", app.ProfileHandler)
	mux.HandleFunc("/profile/upload", app.ProfileUploadHandler)
	mux.HandleFunc("/profile/image", app.ProfileImageHandler)
	mux.HandleFunc("GET /profile/update", app.ProfileUpdateHandler)
	mux.HandleFunc("POST /profile/update", app.ProfileUpdateHandler)
	mux.HandleFunc("DELETE /profile/delete", app.ProfileDeleteHandler)

	//auth
	mux.HandleFunc("/", app.NotFoundHandler) //every unregistered endpoints go here
	mux.HandleFunc("GET /register", app.HandleRegister)
	mux.HandleFunc("POST /register", app.HandleRegister)
	mux.HandleFunc("GET /login", app.HandleLogin)
	mux.HandleFunc("POST /login", app.HandleLogin)
	mux.HandleFunc("POST /logout", app.HandleLogout)
	mux.HandleFunc("GET /{$}", app.HandleHome)
	mux.HandleFunc("GET /forgot-password", app.HandleForgotPassword)
	mux.HandleFunc("POST /forgot-password", app.HandleForgotPassword)
	mux.HandleFunc("GET /reset-password", app.HandleResetPassword)
	mux.HandleFunc("POST /reset-password", app.HandleResetPassword)
	// Public forum routes
	mux.HandleFunc("GET /files/{id}", app.File)
	mux.HandleFunc("GET /posts", app.ListPosts)
	mux.HandleFunc("GET /posts/search", app.SearchPost)
	mux.HandleFunc("GET /posts/{id}", app.GetPostByID)
	mux.HandleFunc("GET /post/{id}", app.PostPage)
	mux.HandleFunc("GET /categories", app.ListCategories)
	mux.HandleFunc("GET /posts/{id}/comments", app.ListComment)

	// Routes that require a valid login session.
	mux.Handle("GET /post/new", app.RequireAuth(app.NewPostForm))
	mux.Handle("POST /post/new", app.RequireAuth(app.CreatePostFromForm))
	mux.Handle("POST /post/{id}/comments", app.RequireAuth(app.CreateCommentFromForm))
	mux.Handle("POST /post/{id}/reaction", app.RequireAuth(app.ReactToPost))
	mux.Handle("POST /comment/{id}/reaction", app.RequireAuth(app.ReactToComment))
	//mux.Handle("POST /posts", app.RequireAuth(app.CreatePost))
	mux.Handle("POST /posts/{id}/comments", app.RequireAuth(app.CreateComment))
}
