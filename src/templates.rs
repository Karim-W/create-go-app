pub const GO_MOD_TEMPLATE: &str = r"module {{module_name}}

go 1.21

";

pub const USECASE_TEMPLATE: &str = r"package usecases

type {{usecase_name_cap}} interface {
}

";

pub const USECASE_IMPL_TEMPLATE: &str = r"package {{usecase_name}}usecase

type {{usecase_name}}_ struct {
}

func New() usecases.{{usecase_name_cap}} {
    return &{{usecase_name}}_{}
}
";

pub const REPOSITORY_TEMPLATE: &str = r"package repositories

type {{repository_name_cap}} interface {
}

";

pub const REPOSITORY_IMPL_TEMPLATE: &str = r"package {{repository_name}}repository

type {{repository_name}}_ struct {
}

func New() repositories.{{repository_name_cap}} {
    return &{{repository_name}}_{}
}
";

pub enum Structures {
    Basic,
    BasicWithFlags,
}

impl Structures {
    #[must_use]
    pub const fn get_template_path(&self) -> &str {
        match self {
            Self::Basic => "basic",
            Self::BasicWithFlags => "bwf",
        }
    }

    #[must_use]
    pub fn resolve_str(template: &str) -> Self {
        match template {
            "basic-cli" => Self::BasicWithFlags,
            _ => Self::Basic,
        }
    }
}
