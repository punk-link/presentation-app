package services

type StaticServer interface {
	Get(hash string) (map[string]any, error)
}
