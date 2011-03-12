package mytools

import "testing"
func TestReverse(t *testing.T){
    i:="ABCD"
    o:=Reverse(i)
    e:="DCBA"
    if o != e {
        t.Error("%q expected %q got %q",i,e,o)
    }
}

func BenchmarkReverse(b *testing.B) {
    s:="ABCD"
    for i:=0;i<b.N;i++{
        Reverse(s)
    }

}
