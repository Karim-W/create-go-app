pub const GO_MOD_TEMPLATE: &str = r#"module {{module_name}}

go 1.20

"#;

pub const USECASE_TEMPLATE: &str = r#"package usecases

type {{usecase_name_cap}} interface {
}

"#;

pub const USECASE_IMPL_TEMPLATE: &str = r#"package {{usecase_name}}usecase

type _{{usecase_name}} struct {
}

func New() usecases.{{usecase_name_cap}} {
    return &_{{usecase_name}}{}
}
"#;

pub const REPOSITORY_TEMPLATE: &str = r#"package repositories

type {{repository_name_cap}} interface {
}

"#;

pub const REPOSITORY_IMPL_TEMPLATE: &str = r#"package {{repository_name}}repository

type _{{repository_name}} struct {
}

func New() repositories.{{repository_name_cap}} {
    return &_{{repository_name}}{}
}
"#;
