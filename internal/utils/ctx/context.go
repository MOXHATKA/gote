package ctx

import "context"

type CustomContext struct {
	GoContext 	context.Context
	Token   	string
}