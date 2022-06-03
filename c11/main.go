package main

// MVC 贫血模型

type Controller struct {
	Svc Service
}

func (c *Controller) GetUserById(userId string) Vo {
	c.Svc.GetUserById(userId)
	// convert Bo to Vo
	return Vo{}
}

type Vo struct {

}

type Service struct {
	Rp Repo
}

func (s *Service) GetUserById(userId string) Bo {
	s.Rp.GetUserById(userId)
	// convert
	return Bo{}
}

type Bo struct {

}

type Repo struct {

}

func (r *Repo) GetUserById(userId string) Entity {
	return Entity{}
}

type Entity struct {
}