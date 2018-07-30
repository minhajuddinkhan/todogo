package models

//GetAllModels GetAllModels
func GetAllModels() []interface{} {

	var models []interface{}
	models = append(models, &Todo{})
	return models
}
