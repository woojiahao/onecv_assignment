package server

type (
	ErrorResponse struct {
		Message string `json:"message"`
	}

	RegisterStudentsDto struct {
		Teacher  string   `json:"teacher" binding:"required,email"`
		Students []string `json:"students" binding:"required,dive,email"`
	}

	CreateTeacherDto struct {
		Email string `json:"email" binding:"required,email"`
	}

	CreateStudentDto struct {
		Email string `json:"email" binding:"required,email"`
	}

	SuspendDto struct {
		Student string `json:"student" binding:"required,email"`
	}

	RetrieveForNotificationsDto struct {
		Teacher      string `json:"teacher" binding:"required,email"`
		Notification string `json:"notification" binding:"required"`
	}
)
