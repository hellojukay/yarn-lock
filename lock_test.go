package lock

import (
	"testing"
)

func Test_New(t *testing.T) {
	l, err := FromFile("./testdata/yarn.lock")
	if err != nil {
		t.Fail()
	}
	if l == nil {
		t.FailNow()
	}
	components := l.Component()
	if len(components) == 0 {
		t.Fail()
	}
	if components[0].Name != "@babel/parser" {
		t.Fatal("prase Name failed")
	}
	if components[0].Integrity != "sha512-7yJPvPV+ESz2IUTPbOL+YkIGyCqOyNIzdguKQuJGnH7bg1WTIifuM21YqokFt/THWh1AkCRn9IgoykTRCBVpzA==" {
		t.Fatal("prase Integrity failed")
	}
	if components[0].Resolved != "https://registry.yarnpkg.com/@babel/parser/-/parser-7.17.3.tgz#b07702b982990bf6fdc1da5049a23fece4c5c3d0" {
		t.Fail()
	}
	if components[0].Version != "7.17.3" {
		t.Fail()
	}
}
