package controllers

type LoginPayload struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
type TokenResponse struct {
	SignedToken  string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
type CreatePostPayload struct {
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}

type UpdatePostPayload struct {
	PostId  uint   `json:"post_id" binding:"required"`
	Title   string `json:"title" binding:"required"`
	Content string `json:"content" binding:"required"`
}
