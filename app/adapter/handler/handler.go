package handler

type Handler struct {
	UserHandler
	// add other handlers below
}

func NewHandler() *Handler {
	return &Handler{
		UserHandler: NewUserHandler(),
		// add other handlers below
	}
}
