use std::{fs, path::Path};

use crate::{models, controllers, utils, packages::find_folder};

pub fn generate_package(path:&str,typ: &str,name: &str) {
    match (typ,name) {
        ("swagger","controllers") => {
            generate_server_from_swagger(path);
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

    for (key, value) in swagger.paths {

        // ==> Get
        if let Some(get_operation) = value.get {

            let mut req = models::RequestHandler::from(&get_operation);

            req.path = Some(key.clone());
            req.method = Some("GET".to_string());
            
            // first tag or default
            let first_tag = req.tags.clone().unwrap_or_else(||vec!["default".to_string()]).first().expect("failed to get route tag").clone();
            
            // check if tag exists
            if !tree.handlers.contains_key(first_tag.clone().as_str()) {
                tree.handlers.insert(first_tag.clone(), vec![]);
            }
            
            // add to tag
            tree.handlers.get_mut(&first_tag.clone()).expect("failed to process route").push(req.clone());
        }

        // ==> Post
        if let Some(post_operation) = value.post {

            let mut req = models::RequestHandler::from(&post_operation);

            req.path = Some(key.clone());
            req.method = Some("POST".to_string());
            
            // first tag or default
            let first_tag = req.tags.clone().unwrap_or_else(||vec!["default".to_string()]).first().expect("failed to get route tag").clone();
            
            // check if tag exists
            if !tree.handlers.contains_key(first_tag.clone().as_str()) {
                tree.handlers.insert(first_tag.clone(), vec![]);
            }

            // add to tag
            tree.handlers.get_mut(&first_tag.clone()).expect("failed to process route").push(req.clone());
        }
        // ==> Put
        if let Some(put_operation) = value.put {

            let mut req = models::RequestHandler::from(&put_operation);

            req.path = Some(key.clone());
            req.method = Some("PUT".to_string());
            
            // first tag or default
            let first_tag = req.tags.clone().unwrap_or_else(||vec!["default".to_string()]).first().expect("failed to get route tag").clone();
            
            // check if tag exists
            if !tree.handlers.contains_key(first_tag.clone().as_str()) {
                tree.handlers.insert(first_tag.clone(), vec![]);
            }
            // add to tag
            tree.handlers.get_mut(&first_tag.clone()).expect("failed to process route").push(req.clone());
        }

        // ==> Delete
        if let Some(delete_operation) = value.delete {

            let mut req = models::RequestHandler::from(&delete_operation);

            req.path = Some(key.clone());
            req.method = Some("DELETE".to_string());

            // first tag or default
            let first_tag = req.tags.clone().unwrap_or_else(||vec!["default".to_string()]).first().expect("failed to get route tag").clone();
            
            // check if tag exists
            if !tree.handlers.contains_key(first_tag.clone().as_str()) {
                tree.handlers.insert(first_tag.clone(), vec![]);
            }
            // add to tag
            tree.handlers.get_mut(&first_tag.clone()).expect("failed to process route").push(req.clone());
        }

        // ==> Patch
        if let Some(patch_operation) = value.patch {

            let mut req = models::RequestHandler::from(&patch_operation);

            req.path = Some(key.clone());
            req.method = Some("PATCH".to_string());

            // first tag or default
            let first_tag = req.tags.clone().unwrap_or_else(||vec!["default".to_string()]).first().expect("failed to get route tag").clone();
            
            // check if tag exists
            if !tree.handlers.contains_key(first_tag.clone().as_str()) {
                tree.handlers.insert(first_tag.clone(), vec![]);
            }
            // add to tag
            tree.handlers.get_mut(&first_tag.clone()).expect("failed to process route").push(req.clone());
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
