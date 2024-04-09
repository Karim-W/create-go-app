pub mod bare;
pub mod cli;
pub mod clients;
pub mod controllers;
pub mod gen;
pub mod models;
pub mod mono;
pub mod openapi;
pub mod packages;
pub mod templates;
pub mod traits;
pub mod utils;

fn main() {
    cli::startup();
}
