package composite

import "testing"

func TestSon_GetParent(t *testing.T) {
	son := Son{}
	GetParentField(&son.P)
}
