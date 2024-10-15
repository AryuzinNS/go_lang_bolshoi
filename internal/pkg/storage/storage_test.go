package storage

// a test file for methods of structure "Storage"
import "testing"

type testCase struct { //testcases for set/get test
	name  string
	key   string
	value string
}

type testCaseFT struct { //testcases for GetKind test
	name    string
	key     string
	value   string
	type_OV string
}

func TestSG(t *testing.T) { // test for set/get procedure
	cases := []testCase{
		{"Hello World", "Hello", "World"},
		{"EmpryTest", "Key2", ""},
		{"Integer test", "key3", "9112001"},
	}
	s, cde := NewStorage()
	if cde != nil {
		t.Errorf("new storage: %v", cde)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)
			tval := s.Get(c.key)
			if *tval != c.value {
				t.Errorf("vals are unequal")
			}
		})
	}
}

func TestTypeOV(t *testing.T) { // test for type of value
	cases := []testCaseFT{
		{"case1", "Key1", "Val1", "S"},
		{"case2", "Key2", "1", "D"},
	}
	s, err := NewStorage()
	if err != nil {
		t.Errorf("new value %v", err)
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			s.Set(c.key, c.value)
			tval := s.GetKind(c.key)

			if tval != c.type_OV {
				t.Errorf("types are not equal")
			}
		})
	}
}
