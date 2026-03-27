package service

type TemplateService struct{}

func NewTemplateService() *TemplateService {
	return &TemplateService{}
}

func (s *TemplateService) GetTemplateList() []string {
	// 先模拟数据(后面接数据库)
	return []string{"template1", "template2", "template3"}
}