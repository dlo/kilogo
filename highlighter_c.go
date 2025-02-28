package main

type CLanguage struct{}

func (l CLanguage) GetHLExtensions() []string {
	return []string{".c", ".h", ".cpp", ".hpp", ".cc"}
}

func (l CLanguage) GetHLKeywords() []string {
	return []string{
		/* C Keywords */
		"auto", "break", "case", "continue", "default", "do", "else", "enum",
		"extern", "for", "goto", "if", "register", "return", "sizeof", "static",
		"struct", "switch", "typedef", "union", "volatile", "while", "NULL",

		/* C++ Keywords */
		"alignas", "alignof", "and", "and_eq", "asm", "bitand", "bitor", "class",
		"compl", "constexpr", "const_cast", "deltype", "delete", "dynamic_cast",
		"explicit", "export", "false", "friend", "inline", "mutable", "namespace",
		"new", "noexcept", "not", "not_eq", "nullptr", "operator", "or", "or_eq",
		"private", "protected", "public", "reinterpret_cast", "static_assert",
		"static_cast", "template", "this", "thread_local", "throw", "true", "try",
		"typeid", "typename", "virtual", "xor", "xor_eq",

		/* C types */
		"int|", "long|", "double|", "float|", "char|", "unsigned|", "signed|",
		"void|", "short|", "auto|", "const|", "bool|",
	}
}
