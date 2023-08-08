package deploy_service

type DeployServiceImpl struct{}

func NewDeployService() DeployService {
	deployService := &DeployServiceImpl{}
	return deployService
}
