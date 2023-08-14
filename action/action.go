package action

type Action struct {
	Name             string
	Method           string
	Parameters       map[string]interface{}
	ParametersConfig map[string]interface{}
	Response         string
	ResponseConfig   map[string]interface{}
	ResponseType     string
}
