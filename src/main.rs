pub mod bare;
pub mod cli;
pub mod gen;
pub mod models;
pub mod utils;
pub mod controllers;
pub mod packages;
pub mod templates;
pub mod traits;

fn main() {
    cli::startup();
}
