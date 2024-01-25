pub mod bare;
pub mod cli;
pub mod clients;
pub mod controllers;
pub mod gen;
pub mod models;
pub mod openapi;
pub mod packages;
pub mod utils;
pub mod templates;
pub mod traits;

fn main() {
    cli::startup();
}
