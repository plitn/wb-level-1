package task21

import "fmt"

/*
Допустим, у нас есть машина, которая готовит кофе.
И тут нам вдруг понадобилось чтобы она готовила еще и чай.
Тогда мы сделаем адаптер, в который будет встраиваться Чай
адаптер реализует интерфейс деланья напитка
теперь Персон может:
-- сделать кофе, если передаст в машину кофе
-- сделать чай, если передаст в машину адаптер чая
*/

type Person struct{}

func (p *Person) makeDrink(maker Machine) {
	fmt.Println("person is making a drink")
	maker.makeDrink()
}

type Machine interface {
	makeDrink()
}

type Coffee struct{}

func (c *Coffee) makeDrink() {
	fmt.Println("making coffee")
}

type Tea struct {
}

func (t *Tea) makeLightDrink() {
	fmt.Println("making tea")
}

type teaAdapter struct {
	teaMaker Tea
}

func (ta *teaAdapter) makeDrink() {
	fmt.Println("converting coffee to tea")
	ta.teaMaker.makeLightDrink()
}
func RunTask21() {
	person := &Person{}
	coffee := &Coffee{}
	tea := &Tea{}
	fluidAdapter := &teaAdapter{
		teaMaker: *tea,
	}
	person.makeDrink(coffee)
	person.makeDrink(fluidAdapter)
}
