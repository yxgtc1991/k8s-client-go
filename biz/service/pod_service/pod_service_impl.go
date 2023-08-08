package pod_service

type PodServiceImpl struct{}

func NewPodService() PodService {
	podService := &PodServiceImpl{}
	return podService
}
