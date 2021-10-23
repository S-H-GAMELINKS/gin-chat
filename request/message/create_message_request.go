package request

type CreateMessageRequest struct {
	Content string `form:"content"`
}
