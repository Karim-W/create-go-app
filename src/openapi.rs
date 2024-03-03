use std::collections::HashMap;

use serde_json::Value;

use crate::{models::Schema, traits::Capitalize};


pub fn build_golang_type(component_tree: HashMap<String,Schema>) -> HashMap<String,String> {
    let mut type_map :HashMap<String,String>= HashMap::new();

    for (key,value) in component_tree{
        type_map.insert(key.clone(),convert_schema_to_go_struct(key.as_str(),&value));
    }

    return type_map;
}


pub fn convert_schema_to_go_struct(name: &str, schema: &Schema) -> String {

    let mut go_struct = format!("type {} struct {{\n", name);

    if schema.clone().items.is_some() {
        let value = schema_to_value(&schema.clone().items.expect("sd"));
        let val = value.as_object().expect("sd");
        return format!("type {name} []{}\n",map_field_type(&val));
    }


    if schema.clone().properties.is_none() {
        return go_struct;
    }

    for (field_name, field_value) in schema.clone().properties.expect("sd") {
        let value = schema_to_value(&field_value.clone());
        let inner_schema = value.as_object().expect("sd");
        let cloned_inner_schema = inner_schema.clone();
        let field_declaration = convert_field_to_go(field_name.as_str(), &cloned_inner_schema);
        go_struct.push_str(&field_declaration);
    }

    go_struct.push_str("}\n");

    go_struct
}

fn schema_to_value(schema: &Schema) -> Value {
    serde_json::to_value(schema).unwrap()
}

fn convert_field_to_go(field_name: &str, field_value: &serde_json::Map<String, Value>) -> String {
    format!("\t{}\t{}\t`json:\"{}\"`\n", field_name.to_string().capitalize_first_letter(), map_field_type(field_value), field_name)
}

fn map_field_type(field_value: &serde_json::Map<String, Value>) -> String {
    if field_value.get("$ref").is_some() {
        if let Some(t) = field_value.get("$ref").and_then(|t| t.as_str()) {
            if t == &Value::Null {
                return "interface{}".to_string();
            }
            return t.split("/").last().unwrap().to_string();
        }
    }

    match field_value.get("type").and_then(|t| t.as_str()) {
        Some(t) => match t {
            "string" => "string".to_string(),
            "integer" => "int".to_string(),
            "number" => "float64".to_string(),
            "boolean" => "bool".to_string(),
            "array" => {
                let mut typ = map_field_type(field_value.get("items").and_then(|t| t.as_object()).unwrap()).clone().to_string().to_owned();
                typ.insert_str(0, "[]");
                typ
            },
            _ =>{
                let v = field_value.get("properties");
                if v.is_some() {
                    let txt= v.expect("sd");
                    if txt == &Value::Null {
                        return "interface{}".to_string();
                    }

                    map_field_type(field_value.get("properties").and_then(|t| t.as_object()).unwrap())

                } else {
                    "interface{}".to_string()
                }

                
            },
        },
        None => {
            match field_value.get("$ref").and_then(|t| t.as_str()) {
                Some(t) => {
                    t.split("/").last().unwrap().to_string()
                },
                None =>{ 
                    let v = field_value.get("properties");

                    if v.is_some() {
                        let txt= v.expect("sd");
                        if txt == &Value::Null {
                            return "interface{}".to_string();
                        }
                        map_field_type(field_value.get("properties").and_then(|t| t.as_object()).unwrap())
                    } else {
                        "interface{}".to_string()
                    }
                },
            }
        }
    }
}




