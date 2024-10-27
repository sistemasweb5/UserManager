package endpoints

type Applicant struct {
	Id           string
	Name         string
	EmailAddress string
	CategoryId   string
}

type ApplicantResponse struct {
	Id           string
	Name         string
	EmailAddress string
}

type Category struct {
	Id  string
	Rol string
}

type WorkSchedule struct {
	Id        string
	StartTime string
	EndTime   string
}

type Specialty struct {
	Id       string
	Name     string
	ClientId string
}

type UserLogin struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type WorkerResponse struct {
	Id           string
	Name         string
	EmailAddress string
	WorkSchedule WorkSchedule
	Specialties  []Specialty
}

type Worker struct {
	Id             string
	Name           string
	EmailAddress   string
	CategoryId     string
	WorkScheduleId string
}

type ApplicantRequest struct {
	Name         string
	EmailAddress string
}

type WorkerRequest struct {
	Name           string
	EmailAddress   string
	WorkScheduleId string
}
