package endpoints

type Client struct {
	Id             string
	Name           string
	EmailAddress   string
	CategoryId     string
	WorkScheduleId string
}

type Category struct {
	Id  string
	Rol string
}

type ClientResponse struct {
	Client       Client
	Category     Category
	WorkSchedule WorkSchedule
	Specialties  []Specialty
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
