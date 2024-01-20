use std::collections::HashMap;

use crate::models;

pub const CONTROLLER_CONSTRUCTOR: &str = r#"package handlers
// Path: .{{CONTROLLERNAME}}.go
//
// This file is generated by the create-go-app tool.
//
import (
    ".{{MODULE}}/cmd/rest"
	".{{MODULE}}/pkg/services/factory"

	"github.com/gin-gonic/gin"
)


type _.{{CONTROLLERNAME}} struct {
    // pls add your usecase here
}

func .{{CONTROLLERNAME}}(
    // pls add your usecase here
) rest.RestHandler[gin.IRouter] {
    return &_.{{CONTROLLERNAME}}{}
}

"#;


pub fn generate_controller(
    module:&str,
    controller_name: &str,
    routes: &Vec<models::RequestHandler>
) -> String {
    let mut buffer  = CONTROLLER_CONSTRUCTOR.replace(".{{CONTROLLERNAME}}", controller_name).replace(".{{MODULE}}", module);
    let mut route_count = 0;
    let mut route_mapping:HashMap<String,(String,String)> = HashMap::new();

    for route in routes {
        route_count += 1;

        let function_name = route.clone().summary.unwrap_or(format!("handler_{}",route_count)).replace(" ", "_").to_lowercase();

        let sanitized_path = route.path.clone().unwrap_or("/".to_string()).replace("{", ":").replace("}", "");

        route_mapping.insert(sanitized_path, (route.method.clone().unwrap_or("GET".to_string()), function_name.clone()));

        buffer.push_str(
            format!("// Handler - {}\nfunc (c *_{}) {}(ctx *gin.Context) {{}}\n",
                    route.clone().description.unwrap_or("".to_string()),
                    controller_name, function_name
        ).as_str());
    }

    buffer.push_str(format!("func (c *_{}) SetupRoutes(rg gin.IRouter) {{\n", controller_name).as_str());
    for (route, (method,function)) in route_mapping {
        buffer.push_str(format!("\trg.{}(\"{}\", c.{})\n",method,route, function).as_str());
    }
    buffer.push_str("}\n");

    return buffer;
}

