use std::fs;

use toml::Value;

pub fn get_service_definition() -> (String, String) {
    // read the service definition from .cga.toml
    let config = fs::read_to_string(".cga.toml").expect("Something went wrong reading the config file");
    // parse the config
    let config: Value = config.parse().unwrap();
    // get module name
    let module_name = config["application"]["module"].as_str().expect("No module name found in .cga.toml file's [application] section");
    // get service name
    let service_name = config["application"]["service"].as_str().expect("No service name found in .cga.toml file's [application] section");
    // return
    (module_name.to_string(), service_name.to_string())
}
