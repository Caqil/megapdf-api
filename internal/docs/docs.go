
package docs

import "github.com/swaggo/swag"

var SwaggerInfo = &swag.Spec{
    Version:          "1.0",
    Host:             "localhost:8080",
    BasePath:         "/api",
    Schemes:          []string{"http"},
    Title:            "REST API",
    Description:      "API documentation for the Go REST API with SQLite",
    InfoInstanceName: "swagger",
    SwaggerTemplate:  "",
}
