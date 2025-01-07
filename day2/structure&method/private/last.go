package private

type Thing struct {
	thing1 string
	thing2 string
	thing3 string
}

func GetThing(thing Thing)[]string{
	output := []string{Thing.thing1,Thing.thing2,Thing.thing3}
	return output 
}


func (t Thing) GetThing1() string{
	return p.thing1
}
func (t Thing) GetThing2() string{
	return p.thing2
}
func (t Thing) GetThing3() string{
	return p.thing3
}

func (t *Thing) SetThing1(t string){
	t.thing1=t
}
func (t *Thing) SetThing2(t string){
	t.thing2=t
}
func (t *Thing) SetThing3(t string){
	t.thing3=t
}