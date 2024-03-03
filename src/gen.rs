use std::{fs, path::Path, collections::HashMap};


use crate::{models::{self, ClientContract, ClientParameter}, controllers, utils, packages::find_folder, traits::Capitalize, openapi::{self, convert_schema_to_go_struct}, clients::CLIENT_CONSTRUCTOR};

pub fn generate_package(path:&str,typ: &str,name: &str) {
    match (typ,name) {
        ("swagger","controllers") => {
            generate_server_from_swagger(path);
        },
        ("swagger","client") => {
            generate_client_from_swagger(path);
        }
        _ => {
            println!("Unknown type: {typ}");
        }
    }
}



fn generate_server_from_swagger(path:&str) {

    let content = fs::read_to_string(path).expect("Something went wrong reading the file");

    let swagger: models::OpenApi = serde_json::from_str(&content).expect("failed to parse swagger file");
    
    let mut tree = models::RequestTree::new();

    let (module_name, _) = utils::get_service_definition();

    for (key, value) in swagger.paths.iter() {

        // ==> Get
        if let Some(get_operation) = &value.get {
            handle_swagger_server_route("GET",get_operation, &key, &mut tree);
        }

        // ==> Post
        if let Some(post_operation) = &value.post {
            handle_swagger_server_route("POST",post_operation, &key, &mut tree);
        }
        // ==> Put
        if let Some(put_operation) = &value.put {
            handle_swagger_server_route("PUT",put_operation, &key, &mut tree);
        }

        // ==> Delete
        if let Some(delete_operation) = &value.delete {
            handle_swagger_server_route("DELETE",delete_operation, &key, &mut tree);
        }

        // ==> Patch
        if let Some(patch_operation) = &value.patch {
            handle_swagger_server_route("PATCH",patch_operation, &key, &mut tree);
        }
    }

    let mut path_string = find_folder("handlers");

    if path_string.is_none() {
        let path = find_folder("rest");

        if path.is_none() {
            println!("cannot find parent package");
            return;
        }

        // create the handlers folder
        let path = path.expect("cannot find parent package");
        let path = Path::new(&path);
        let path = path.join("handlers");

        fs::create_dir_all(path.clone()).expect("cannot create handlers folder");

        path_string = Some(path.to_str().expect("failed to realize path").to_string());
    }

    let path_string = path_string.expect("cannot find parent package");

    for (handler,routes) in tree.handlers {

        let path = path_string.clone();
        let path = Path::new(&path);
        //create file with extension .go
        let path = path.join(handler.to_string() + ".go");

        let buff = controllers::generate_controller(&module_name,handler.clone().as_str(),&routes);
        fs::write(path, buff).expect("Unable to write file");
    }
}

fn handle_swagger_server_route(method:&str ,op: &models::Operation, key: &String, tree: &mut models::RequestTree) {

    let mut req = models::RequestHandler::from(op);

    req.path = Some(key.clone());
    req.method = Some(method.to_string());

    // first tag or default
    let first_tag = req.tags.clone().unwrap_or_else(||vec!["default".to_string()]).first().expect("failed to get route tag").clone();
            
    // check if tag exists
    if !tree.handlers.contains_key(first_tag.clone().as_str()) {
        tree.handlers.insert(first_tag.clone(), vec![]);
    }
    // add to tag
    tree.handlers.get_mut(&first_tag.clone()).expect("failed to process route").push(req.clone());
}



fn generate_client_from_swagger(path:&str) {

    let content = fs::read_to_string(path).expect("Something went wrong reading the file");

    let swagger: models::OpenApi = serde_json::from_str(&content).expect("failed to parse swagger file");

    let mut generated_contract_models = openapi::build_golang_type(swagger.components.expect("failed to get components").schemas.expect("failed to get schemas"));

    // [schema reference]:(type_name, type_definition)
    let mut schemes :HashMap<String, ClientContract> = HashMap::new();

    let package_name = swagger.info.title.to_lowercase();

    let (module_name, _) = utils::get_service_definition();

    for (key, value) in swagger.paths.iter() {

        let key = key.as_str();

        // ==> Get
        if let Some(get_operation) = &value.get {
            gen_client("Get",get_operation, key, &mut generated_contract_models, &mut schemes);
        }

        // ==> Post
        if let Some(post_operation) = &value.post {
            gen_client("Post",post_operation, key, &mut generated_contract_models, &mut schemes);
        }

        // ==> Put
        if let Some(put_operation) = &value.put {
            gen_client("Put",put_operation, key, &mut generated_contract_models, &mut schemes);
        }

        // ==> Delete
        if let Some(delete_operation) = &value.delete {
            gen_client("Delete",delete_operation, key, &mut generated_contract_models, &mut schemes);
        }

        // ==> Patch
        if let Some(patch_operation) = &value.patch {
            gen_client("Patch",patch_operation, key, &mut generated_contract_models, &mut schemes);
        }
    }


    let models_buffer :String =
        format!("package {package_name}\n\n{}",generated_contract_models.clone().into_values().collect::<Vec<String>>().join("\n"));

    let  path = find_folder(package_name.clone().as_str());

    if path.is_none() {
        let path = find_folder("adapters");

        if path.is_none() {
            println!("cannot find parent package");
            return;
        }

        // create the handlers folder
        let path = path.expect("cannot find parent package");
        let path = Path::new(&path);
        let path = path.join(package_name.clone());

        fs::create_dir_all(path.clone()).expect("cannot create package folder");

    }

    let path = path.expect("cannot find parent package");

    let path = Path::new(&path);

    // create the handlers folder
    let models_path = path.join("models.go");

    fs::write(models_path, models_buffer).expect("Unable to write file");

    

    for (key,value) in schemes.clone() {
        let file_path = path.join(key.clone() + ".go");

        fs::write(file_path, value.convert_to_go().replace("{{.CLIENT}}", &package_name).replace("{{.MODULE}}", &module_name)).expect("Unable to write file");
    }

    // constuctors
    let sigs :Vec<String> = schemes.clone().into_values().map(|x| x.get_signature()).collect();

    let buff = CLIENT_CONSTRUCTOR.replace("{{.CLIENT}}", &package_name).replace("{{.MODULE}}", &module_name).replace("{{.FUNCS}}", &sigs.join("\n"));

    let file_path = path.join("client.go");

    fs::write(file_path, buff).expect("Unable to write file");
}

fn gen_client(
    method:&str,
    operation: &models::Operation,
    key: &str, 
    generated_contract_models: &mut HashMap<String, String>, 
    schemes: &mut HashMap<String, ClientContract>,
    ) {

    let op = &operation.clone();

    let name = &op.operation_id.as_ref().expect("failed to get operation id").capitalize_first_letter();

    let path = key;

    let method = method.to_string();

    let description = op.description.as_ref().expect("failed to get description").clone();

    let mut return_type = String::new();

    let responses = op.responses.as_ref().expect(format!("failed to get responses for {method} {key}").as_str());

    let mut args :Vec<ClientParameter> = vec![];

    let mut body :Option<String>= None;

    let op = &operation.clone();

    if op.parameters.is_some() {
        for param in op.parameters.as_ref().expect("failed to get parameters") {
            let arg = ClientParameter::from_parameter(&param);
            args.push(arg);
        }
    }


    // handle request body
    if operation.request_body.is_some(){

        let content = op.request_body.as_ref().expect("failed to get request_body").content
            .get("application/json").expect("failed to get content for JSON response")
            .schema.clone().expect("failed to get JSON Schema");

        if let Some(schema_path) = content.r#ref {
            let schema_type_name = schema_path.split("/").last().expect("failed to get schema name");

            body = Some(schema_type_name.to_string());
        }else{

            let schema = content.clone();

            let struct_name = format!("{}Request", name.clone());
            
             body = Some(struct_name.clone());

            let def = convert_schema_to_go_struct(struct_name.as_str(), &schema);
        
            generated_contract_models.insert(struct_name.clone(), def.clone());
        }


    }

    // will only support one successful response
    for (status, resp) in responses.iter(){
    
        //check if status is not 2xx
        if !status.starts_with("2") {
            continue;
        }

        if resp.content.is_none() {
            continue;
        }

        if resp.clone()
            .content.expect("failed to get content for response")
            .get("application/json")
            .is_none() {

            continue;
        }

        // will only support JSON response for now
        let content = resp.content.as_ref().expect("failed to get content for response")
            .get("application/json").expect("failed to get content for JSON response")
            .schema.clone().expect("failed to get JSON Schema");

        if let Some(schema_path) = content.r#ref {
            let schema_type_name = schema_path.split("/").last().expect("failed to get schema name");

            return_type = schema_type_name.to_string();
            continue;
        }

        let schema = content.clone();

        let struct_name = format!("{}Result", name.clone());
        return_type = struct_name.clone();

        let def = convert_schema_to_go_struct(struct_name.as_str(), &schema);
    
        generated_contract_models.insert(struct_name.clone(), def.clone());
    }


    schemes.insert(name.clone(), ClientContract{
        description: Some(description.clone()),
        args,
        body,
        return_type,
        path: path.to_string(),
        method,
        name: name.clone(),
    });
}

