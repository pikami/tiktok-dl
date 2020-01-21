package unittestutil

// AssertString - Check if two strings match
func (tu *TestUtil) AssertString(actual string, expected string, name string) {
	if actual != expected {
		tu.T.Errorf("%s is incorrect: Expected '%s', but got '%s'", name, expected, actual)
	}
}

// AssertInt - Check if two intagers match
func (tu *TestUtil) AssertInt(actual int, expected int, name string) {
	if actual != expected {
		tu.T.Errorf("%s is incorrect: Expected '%d', but got '%d'", name, expected, actual)
	}
}
