package v1

type Configuration struct{}

func NewConfiguration() *Configuration {
	return &Configuration{}
}

func (*Configuration) InitUserService() *UserService {
	return NewUserService()
}
