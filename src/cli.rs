use clap::{arg, Command};

use crate::{gen::generate_package, packages, templates};

fn base_command() -> Command {
    Command::new("create-go-app")
        .version("0.1.7")
        .about("Create a new Go application")
        .author("Karim-w")
        .arg_required_else_help(false)
        .allow_external_subcommands(true)
}

fn add_new_sub_command(cmd: Command) -> Command {
    cmd.subcommand(
        Command::new("new")
            .about("create a new package")
            .arg(arg!(<typ> "The type of package to create"))
            .arg(arg!(<name> "The name of the package"))
            .arg_required_else_help(true),
    )
}

fn add_template_sub_command(cmd: Command) -> Command {
    cmd.subcommand(
        Command::new("template")
            .about("create app using template")
            .arg(arg!(<name> "the name of the template to use"))
            .arg_required_else_help(true),
    )
}

fn add_mono_repo_sub_command(cmd: Command) -> Command {
    cmd.subcommand(
        Command::new("mono")
            .about("create a new monorepo")
            .arg(arg!(<typ> "The type of package to create"))
            .arg_required_else_help(true),
    )
}

#[allow(clippy::cognitive_complexity)]
fn add_gen_sub_command(cmd: Command) -> Command {
    cmd.subcommand(
        Command::new("gen")
            .about("generate a new package")
            .arg(arg!(<typ> "The type of generator to use"))
            .arg(arg!(<path> "The path of definition file"))
            .arg(arg!(<name> "The name of the package to generate"))
            .arg_required_else_help(true),
    )
}

fn build_command() -> Command {
    let mut cmd = base_command();
    cmd = add_new_sub_command(cmd);
    cmd = add_gen_sub_command(cmd);
    cmd = add_template_sub_command(cmd);
    cmd = add_mono_repo_sub_command(cmd);

    cmd
}

/// # Panics
pub fn startup() {
    let cmd = build_command();

    let matches = cmd.get_matches();

    match matches.subcommand() {
        Some(("new", sub_matches)) => {
            let typ = sub_matches
                .get_one::<String>("typ")
                .expect("No type provided");
            let name = sub_matches
                .get_one::<String>("name")
                .expect("No name provided");
            packages::handle_add(typ.as_str(), name.as_str());
        }
        Some(("gen", sub_matches)) => {
            let typ = sub_matches
                .get_one::<String>("typ")
                .expect("No type provided");
            let path = sub_matches
                .get_one::<String>("path")
                .expect("No path provided");
            let name = sub_matches
                .get_one::<String>("name")
                .expect("No name provided");

            generate_package(path, typ, name);
        }
        Some(("template", sub_matches)) => {
            let name = sub_matches
                .get_one::<String>("name")
                .expect("No name provided");
            println!("Creating Service: {name}");
            crate::bare::set_up_basic_app(
                templates::Structures::resolve_str(name).get_template_path(),
            );
        }
        Some(("mono", sub_matches)) => {
            let alt: String = "app".to_string();
            let typ = sub_matches.get_one::<String>("typ").unwrap_or(&alt);
            match typ.as_str() {
                "app" => {
                    crate::mono::setup_monorepo_application();
                }
                _ => {
                    crate::mono::setup_monorepo();
                }
            }
        }
        _ => {
            crate::bare::set_up_basic_app("basic");
        }
    }
}
