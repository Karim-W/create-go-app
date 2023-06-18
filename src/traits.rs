pub trait Capitalize {
    fn capitalize_first_letter(&self) -> String;
}

impl Capitalize for String {
    fn capitalize_first_letter(&self) -> String {
        let mut chars = self.chars();
        match chars.next() {
            None => String::new(),
            Some(f) => f.to_uppercase().collect::<String>() + chars.as_str(),
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
