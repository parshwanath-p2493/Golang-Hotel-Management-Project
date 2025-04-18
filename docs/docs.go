// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": ["http"],
    openapi: "3.0.3",
    "info": {
        "description": "This is a Backend Hotel Management Project.",
        "title": "Hotel Management ",
        "contact": {},
        "version": "1.0"
    },
    "host": "localhost:8080",
    "basePath": "/",
   
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:2493",
	BasePath:         "apii/v1",
	Schemes:          []string{"http"},
	Title:            "Hotel Management",
	Description:      "This is a Backend Hotel Management Project.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
