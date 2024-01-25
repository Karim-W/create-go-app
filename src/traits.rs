pub trait Capitalize {
    fn capitalize_first_letter(&self) -> String;
}

impl Capitalize for String {
    #[allow(clippy::option_if_let_else,clippy::needless_return)]
    fn capitalize_first_letter(&self) -> String {
        let mut chars = self.chars();

        match chars.next() {
            None => Self::new(),
            Some(f) => f.to_uppercase().collect::<Self>() + chars.as_str(),
        }

    }
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_capitalize() {
        let s = String::from("hello");
        assert_eq!(s.capitalize_first_letter(), "Hello");
    }

    #[test]
    fn test_capitalize_empty() {
        let s = String::from("");
        assert_eq!(s.capitalize_first_letter(), "");
    }

    #[test]
    fn test_capitalize_one_letter() {
        let s = String::from("h");
        assert_eq!(s.capitalize_first_letter(), "H");
    }

    #[test]
    fn test_capitalize_one_letter_uppercase() {
        let s = String::from("H");
        assert_eq!(s.capitalize_first_letter(), "H");
    }
}
