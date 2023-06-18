use std::{
    fs,
    io::{Read, Write},
    path::Path,
};

use walkdir::WalkDir;

use crate::templates;

pub fn set_up_basic_app() {
    println!("Please Enter The Service Name: ");
    let mut service_name = String::new();
    std::io::stdin().read_line(&mut service_name).unwrap();
    let service_name = service_name.trim().to_lowercase();
    let mut module_name = service_name.clone();
    println!("would you like to prefix the service name? (y/n)");
    let mut confirm = String::new();
    std::io::stdin().read_line(&mut confirm).unwrap();
    let confirm = confirm.trim().to_lowercase();
    if confirm == "y" {
        println!("Please Enter The Prefix: ");
        let mut prefix = String::new();
        std::io::stdin().read_line(&mut prefix).unwrap();
        let prefix = prefix.trim().to_lowercase();
        module_name = format!("{}/{}", prefix, service_name);
    }
    println!("Creating Service: {} using default template", service_name);
    // check if directory is empty
    let cmd = "https://github.com/Karim-W/create-go-app.git".to_string();
    let output = std::process::Command::new("git")
        .args(&["clone", cmd.as_str()])
        .output()
        .expect("failed to execute process");
    if !output.status.success() {
        println!(
            "failed to execute process: {}",
            String::from_utf8_lossy(&output.stderr)
        );
        std::process::exit(1);
    }
    println!("{}", String::from_utf8_lossy(&output.stderr));
    let output = std::process::Command::new("mv")
        .args(&[
            "./create-go-app/examples/basic",
            format!("./{}", service_name).as_str(),
        ])
        .output()
        .expect("failed to execute process");
    if !output.status.success() {
        println!(
            "failed to execute process: {}",
            String::from_utf8_lossy(&output.stderr)
        );
        std::process::exit(1);
    }
    let output = std::process::Command::new("rm")
        .args(&["-rf", "create-go-app"])
        .output()
        .expect("failed to execute process");
    if !output.status.success() {
        println!(
            "failed to execute process: {}",
            String::from_utf8_lossy(&output.stderr)
        );
        std::process::exit(1);
    }
    println!("{}", String::from_utf8_lossy(&output.stderr));
    let go_mod = templates::GO_MOD_TEMPLATE.replace("{{module_name}}", module_name.as_str());
    std::fs::write(format!("./{}/go.mod", service_name), go_mod).unwrap();
    let dir_path = format!("./{}", service_name);
    // walk through the directory and replace the template files with the service name
    // and module name
    for entry in WalkDir::new(dir_path).follow_links(true) {
        if entry.is_err() {
            continue;
        }
        let entry = entry.unwrap();
        let file_path = entry.path();
        if file_path.is_file() {
            replace_handlebars_in_file(file_path, service_name.clone(), module_name.clone())
                .expect("Failed to replace handlebars in file");
        }
    }
    println!("Service Created Successfully");
    println!("Running go mod tidy");
    // move into the directory and run go mod tidy
    std::env::set_current_dir(format!("./{}", service_name)).unwrap();
    let output = std::process::Command::new("go")
        .args(&["mod", "tidy"])
        .output()
        .expect("failed to execute process");
    if !output.status.success() {
        println!(
            "failed to execute process: {}",
            String::from_utf8_lossy(&output.stderr)
        );
        std::process::exit(1);
    }
    println!("{}", String::from_utf8_lossy(&output.stderr));
    println!("go mod tidy completed");
    println!("Service Created Successfully");
    println!("would you like to initailize a git repo? (y/n)");
    let mut confirm = String::new();
    std::io::stdin().read_line(&mut confirm).unwrap();
    let confirm = confirm.trim().to_lowercase();
    if confirm == "y" {
        let output = std::process::Command::new("git")
            .args(&["init"])
            .output()
            .expect("failed to execute process");
        if !output.status.success() {
            println!(
                "failed to execute process: {}",
                String::from_utf8_lossy(&output.stderr)
            );
            std::process::exit(1);
        }
        println!("{}", String::from_utf8_lossy(&output.stderr));
        println!("git repo initialized");
    }
    println!("done <3");
}

fn replace_handlebars_in_file(
    file_path: &Path,
    service_name: String,
    module: String,
) -> std::io::Result<()> {
    let mut file = fs::File::open(file_path)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;

    // Perform the desired replacement
    let modified_contents = contents.replace("{{.moduleName}}", module.as_str());
    let modified_contents = modified_contents.replace("{{.serviceName}}", service_name.as_str());

    let mut file = fs::File::create(file_path)?;
    file.write_all(modified_contents.as_bytes())?;
    Ok(())
}
