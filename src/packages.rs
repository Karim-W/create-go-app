use std::{fs, io::Write, path::Path};

use walkdir::WalkDir;

use crate::{templates, traits::Capitalize};

#[derive(Debug, Clone, Copy)]
pub enum Packages {
    Repository,
    Service,
    Usecase,
    Handler,
    Adapter,
}

impl Packages {
    #[must_use]
    pub const fn get_name(&self) -> &str {
        match self {
            Self::Repository => "repository",
            Self::Service => "service",
            Self::Usecase => "usecase",
            Self::Handler => "handler",
            Self::Adapter => "adapter",
        }
    }

    #[must_use]
    pub fn from_string(s: &str) -> Option<Self> {
        match s.to_lowercase().as_str() {
            "repository" => Some(Self::Repository),
            "service" => Some(Self::Service),
            "usecase" => Some(Self::Usecase),
            "handler" => Some(Self::Handler),
            "adapter" => Some(Self::Adapter),
            _ => None,
        }
    }
}

pub fn handle_add(typ: &str, name: &str) {

    Packages::from_string(typ).map_or_else(||{
        println!("Package type not found");
    }, |p|  
    match p {
        Packages::Repository => {
            add_repository(name);
        }
        Packages::Service => {
            add_service(name);
        }
        Packages::Usecase => {
            add_usecase(name);
        }
        Packages::Handler => {
            add_handler(name);
        }
        Packages::Adapter => {
            add_adapter(name);
        }
    }
    );
}

/// # Panics
pub fn add_repository(name: &str) {

    println!("Adding repository: {name}");

    let mut path_string = find_folder("repositories");

    if path_string.is_none() {
        let path = find_folder("pkg");

        if path.is_none() {
            println!("cannot find parent package");
            return;
        }

        // create the usecases folder
        let path = path.expect("failed to get path");
        let path = Path::new(&path);
        let path = path.join("repositories");

        fs::create_dir_all(path.clone()).expect("failed to create repositories folder");

        path_string = Some(path.to_str().expect("failed to get path").to_string());
    }

    let path = path_string.expect("failed to get path");

    let uc_path = Path::new(&path);

    //create file with extension .go
    let path = uc_path.join(name.to_string() + ".go");

    let file = fs::File::create(path).expect("failed to create file");

    let mut writer = std::io::BufWriter::new(file);

    let repo_name = name.to_string();

    let repo_name = repo_name.replace(['-',' ','.'], "_");

    let content = templates::REPOSITORY_TEMPLATE.replace("{{repository_name}}", &repo_name);

    let content = content.replace(
        "{{repository_name_cap}}",
        &repo_name.capitalize_first_letter(),
    );

    writer.write_all(content.as_bytes()).expect("failed to write file");

    // create the usecase implementation
    let path = uc_path.join(repo_name.clone());

    fs::create_dir_all(path.clone()).expect("failed to create usecase folder");

    let path = path.join("repo.go");

    let file = fs::File::create(path).expect("failed to create file");

    let mut writer = std::io::BufWriter::new(file);

    let content = templates::REPOSITORY_IMPL_TEMPLATE.replace("{{repository_name}}", &repo_name);

    let content = content.replace(
        "{{repository_name_cap}}",
        &repo_name.capitalize_first_letter(),
    );

    writer.write_all(content.as_bytes()).expect("failed to write file");
}

/// # Panics
pub fn add_service(name: &str) {

    println!("Adding service: {name}");

    let mut path_string = find_folder("services");

    if path_string.is_none() {

        let path = find_folder("pkg");

        if path.is_none() {
            println!("cannot find parent package");
            return;
        }

        let path = path.expect("failed to get path");
        let path = Path::new(&path);
        let path = path.join("services");

        fs::create_dir_all(path.clone()).expect("failed to create services folder");

        path_string = Some(path.to_str().expect("failed to get path").to_string());
    }

    let path = path_string.expect("failed to get path");

    let path = Path::new(&path);

    let path = path.join(name.to_string() + ".go");

    fs::File::create(path).expect("failed to create file");
}

/// # Panics
pub fn add_usecase(name: &str) {
    println!("Adding usecase: {name}");

    let mut path_string = find_folder("usecases");

    if path_string.is_none() {

        let path = find_folder("pkg");

        if path.is_none() {
            println!("cannot find parent package");
            return;
        }

        // create the usecases folder
        let path = path.expect("failed to get path");
        let path = Path::new(&path);
        let path = path.join("usecases");

        fs::create_dir_all(path.clone()).expect("failed to create usecases folder");

        path_string = Some(path.to_str().expect("failed to create path").to_string());
    }

    let path = path_string.expect("failed to get path");

    let uc_path = Path::new(&path);

    //create file with extension .go
    let path = uc_path.join(name.to_string() + ".go");

    let file = fs::File::create(path).expect("failed to create file");

    let mut writer = std::io::BufWriter::new(file);

    let usecase_name = name.to_string();

    let usecase_name = usecase_name.replace(['-',' ','.'], "_");

    let content = templates::USECASE_TEMPLATE.replace("{{usecase_name}}", &usecase_name);

    let content = content.replace(
        "{{usecase_name_cap}}",
        &usecase_name.capitalize_first_letter(),
    );

    writer.write_all(content.as_bytes()).expect("failed to write file");

    // create the usecase implementation
    let path = uc_path.join(usecase_name.clone());

    fs::create_dir_all(path.clone()).expect("failed to create folder");

    let path = path.join("usecase.go");

    let file = fs::File::create(path).expect("failed to create file");

    let mut writer = std::io::BufWriter::new(file);

    let content = templates::USECASE_IMPL_TEMPLATE.replace("{{usecase_name}}", &usecase_name);

    let content = content.replace(
        "{{usecase_name_cap}}",
        &usecase_name.capitalize_first_letter(),
    );

    writer.write_all(content.as_bytes()).expect("failed to write file");
}

/// # Panics
pub fn add_handler(name: &str) {
    println!("Adding handler: {name}");

    let mut path_string = find_folder("handlers");

    if  path_string.is_none() {

        let path = find_folder("rest");

        if path.is_none() {
            println!("cannot find parent package");
            return;
        }

        // create the handlers folder
        let path = path.expect("failed to get path");
        let path = Path::new(&path);

        let path = path.join("handlers");

        fs::create_dir_all(path.clone()).expect("failed to create folder");

        path_string = Some(path.to_str().expect("failed to get path").to_string());
    }

    let path = path_string.expect("failed to get path");
    let path = Path::new(&path);

    //create file with extension .go
    let path = path.join(name.to_string() + ".go");

    fs::File::create(path).expect("failed to create file");
}

#[allow(clippy::needless_return)]
#[must_use]
/// # Panics
pub fn find_folder(name: &str) -> Option<String> {
    let path = Path::new(".");

    for entry in WalkDir::new(path).follow_links(true) {

        // check if entry is valid, if not continue
        if entry.is_err() {
            continue;
        }

        let entry = entry.expect("failed to get entry");

        if entry.file_type().is_dir() && entry.file_name().to_str().expect("failed to find") == name {
             let path_found = entry.path().to_str().expect("failed to realize path").to_string();

             return Some(path_found);
        }

        continue;
    }

     return None;
}

/// # Panics
pub fn add_adapter(name: &str) {
    // check if adapter is supported
    match name {
        "posty" | "rdb" | "rabbit" |"id"=> {
            println!("Adding adapter: {name}");
        }
        _ => {
            println!("{name} is not a supported adapter");
            return;
        }
    }
    // check if folder exists
    let mut path_string = find_folder("adapters");

    if path_string.is_none() {

        let path = find_folder("pkg");

        if  path.is_none() {
            println!("cannot find parent package");
            return;
        }

        // create the adapter folder
        let path = path.expect("failed to realize path");
        let path = Path::new(&path);
        let path = path.join("adapters");

        fs::create_dir_all(path.clone()).expect("failed to create adapters folder");
        path_string = Some(path.to_str().expect("failed to realize path").to_string());
    }
    
    if path_string.is_none() {
        println!("cannot find parent package");
        return;
    }

    let path = path_string.expect("failed to realize path");
    let adapter_path = Path::new(&path);
    // checkout the file
    let cmd = "https://github.com/Karim-W/create-go-app.git".to_string();

    let output = std::process::Command::new("git")
        .args(["clone", cmd.as_str(), ".cga-temp"])
        .output()
        .expect("failed to execute process");

    if !output.status.success() {
        println!(
            "failed to execute process: {}",
            String::from_utf8_lossy(&output.stderr)
        );
        std::process::exit(1);
    }

    let output = std::process::Command::new("mv")
        .args(&[
            format!("./.cga-temp/examples/adapters/{name}"),
            adapter_path
                .to_str()
                .expect("failed to realize path")
                .to_string(),
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
        .args(["-rf", ".cga-temp"])
        .output()
        .expect("failed to execute process");

    if !output.status.success() {
        println!(
            "failed to execute process: {}",
            String::from_utf8_lossy(&output.stderr)
        );
        std::process::exit(1);
    }
}
