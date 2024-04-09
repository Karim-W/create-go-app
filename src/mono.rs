use std::{
    fs,
    io::{Read, Write},
    path::Path,
};

use walkdir::WalkDir;

use crate::utils;

#[allow(clippy::too_many_lines)]
/// # Panics
pub fn setup_monorepo() {
    println!("Please Enter The Service Name: ");

    let mut service_name = String::new();

    std::io::stdin()
        .read_line(&mut service_name)
        .expect("failed to read line");

    let service_name = service_name.trim().to_lowercase();

    // let mut module_name = service_name.clone();
    let mut module_name = String::new();

    // i just wanna shut the compiler up or clippy w/e
    if module_name.contains('/') {
        println!("module name cannot contain '/'");
        std::process::exit(1);
    }

    println!("would you like to prefix the service name? (y/n)");

    let mut confirm = String::new();

    std::io::stdin()
        .read_line(&mut confirm)
        .expect("failed to read line");

    let confirm = confirm.trim().to_lowercase();

    if confirm == "y" {
        println!("Please Enter The Prefix: ");

        let mut prefix = String::new();
        std::io::stdin()
            .read_line(&mut prefix)
            .expect("failed to read line");

        let prefix = prefix.trim().to_lowercase();

        module_name = format!("{prefix}/{service_name}");
    } else {
        module_name = service_name.clone();
    }

    println!("Creating Service: {service_name}");

    // check if directory is empty
    let cmd = "https://github.com/Karim-W/create-go-app.git".to_string();

    let output = std::process::Command::new("git")
        .args(["clone", cmd.as_str()])
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
        .args([
            format!("./create-go-app/examples/monorepo").as_str(),
            format!("./{service_name}").as_str(),
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
        .args(["-rf", "create-go-app"])
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

    let dir_path = format!("./{service_name}");
    // walk through the directory and replace the template files with the service name
    // and module name
    for entry in WalkDir::new(dir_path).follow_links(true) {
        if entry.is_err() {
            continue;
        }

        let entry = entry.expect("Failed to get entry");

        let file_path = entry.path();

        if file_path.is_file() {
            replace_handlebars_in_file(file_path, service_name.as_str(), module_name.as_str());
        }
    }

    println!("Service Created Successfully");
    println!("Running go mod tidy");

    // move into the directory and run go mod tidy
    std::env::set_current_dir(format!("./{service_name}")).expect("Failed to change directory");

    let output = std::process::Command::new("go")
        .args(["work", "sync"])
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

    println!("go workspace completed");
    println!("Service Created Successfully");
    println!("would you like to initailize a git repo? (y/n)");

    let mut confirm = String::new();

    std::io::stdin()
        .read_line(&mut confirm)
        .expect("failed to read line");

    let confirm = confirm.trim().to_lowercase();

    if confirm == "y" {
        let output = std::process::Command::new("git")
            .args(["init"])
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
    service_name: &str,
    module: &str,
) -> std::io::Result<()> {
    let mut file = fs::File::open(file_path)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;

    // Perform the desired replacement
    let modified_contents = contents.replace("{{.moduleName}}", module);
    let modified_contents = modified_contents.replace("{{.serviceName}}", service_name);

    let mut file = fs::File::create(file_path)?;
    file.write_all(modified_contents.as_bytes())?;

    Ok(())
}

pub fn setup_monorepo_application() {
    println!("Please Enter The application Name: ");

    let mut service_name = String::new();

    std::io::stdin()
        .read_line(&mut service_name)
        .expect("failed to read line");

    let service_name = service_name.trim().to_lowercase();

    let (module_name, _) = utils::get_service_definition();

    println!("Adding application: {service_name}");

    // check if directory is empty
    let cmd = "https://github.com/Karim-W/create-go-app.git".to_string();

    let output = std::process::Command::new("git")
        .args(["clone", cmd.as_str()])
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
        .args([
            format!("./create-go-app/examples/monoapp").as_str(),
            format!("./apps/{service_name}").as_str(),
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
        .args(["-rf", "create-go-app"])
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

    let dir_path = format!("./apps/{service_name}");
    // walk through the directory and replace the template files with the service name
    // and module name
    for entry in WalkDir::new(dir_path).follow_links(true) {
        if entry.is_err() {
            continue;
        }

        let entry = entry.expect("Failed to get entry");

        let file_path = entry.path();

        if file_path.is_file() {
            let _ =
                replace_handlebars_in_file(file_path, service_name.as_str(), module_name.as_str());
        }
    }

    println!("Application Added Successfully");
    println!("Adding application to workspace");
    let output = std::process::Command::new("go")
        .args(["work", "use", format!("./apps/{service_name}").as_str()])
        .output()
        .expect("failed to execute process");

    if !output.status.success() {
        println!(
            "failed to execute process: {}",
            String::from_utf8_lossy(&output.stderr)
        );
        std::process::exit(1);
    }

    println!("syncing go workspace");

    let output = std::process::Command::new("go")
        .args(["work", "sync"])
        .output()
        .expect("failed to execute process");

    if !output.status.success() {
        println!(
            "failed to execute process: {}",
            String::from_utf8_lossy(&output.stderr)
        );
        std::process::exit(1);
    }

    println!("go workspace completed");

    println!("running go mod tidy");

    // move into the directory and run go mod tidy
    std::env::set_current_dir(format!("./apps/{service_name}"))
        .expect("Failed to change directory");

    let output = std::process::Command::new("go")
        .args(["work", "sync"])
        .output()
        .expect("failed to execute process");

    if !output.status.success() {
        println!(
            "failed to execute process: {}",
            String::from_utf8_lossy(&output.stderr)
        );
        std::process::exit(1);
    }

    println!("done <3");
}
