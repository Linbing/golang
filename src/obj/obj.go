package obj

import "log"

type Obj struct {
	Name string
	Age  int
}

func (o *Obj) ModifyAge(age int) {
	o.Age = age
}

func (o Obj) Notify() error {
	log.Printf("User : name is %s ...Age is %d ...\n", o.Name, o.Age)
	return nil
}
