use clap::{arg, Command};

use crate::packages;

pub fn startup() {
    let cmd = Command::new("create-go-app")
        .version("0.1.2")
        .about("Create a new Go application")
        .author("Karim-w")
        .arg_required_else_help(false)
        .allow_external_subcommands(true)
        .subcommand(
            Command::new("new")
                .about("create a new package")
                .arg(arg!(<typ> "The type of package to create"))
                .arg(arg!(<name> "The name of the package"))
                .arg_required_else_help(true),
        );

    let matches = cmd.get_matches();
    match matches.subcommand() {
        Some(("new", sub_matches)) => {
            let typ = sub_matches.get_one::<String>("typ").unwrap();
            let name = sub_matches.get_one::<String>("name").unwrap();
            packages::handle_add(typ.as_str(), name.as_str());
        }
        _ => {
            crate::bare::set_up_basic_app();
        }
    }
}
