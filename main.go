package main

import (
	"fmt"
	"reflect"

	"github.com/antonmedv/expr"
)

type User struct {
	Login        string
	Name         string
	Email        string
	privateNotes string
	salary       int
}

func getUsers() []User {

	users := []User{
		{
			Login:        "katelynndel",
			Name:         "Tirzah Menges",
			Email:        "elicia_lewteree1@alternative.fbu",
			privateNotes: "Township ping marvel chassis date sorry screensaver, paintings nelson letting lip employ enemy tasks, continues viewing shapes warranty. ",
			salary:       1,
		},
		{
			Login:        "shellycjpk",
			Name:         "Aryana Renn",
			Email:        "crosby_swiftb@casting.gd",
			privateNotes: "Controversy educational workshop website waters habitat decline, nightmare argue departure eye shapes humidity chelsea, settled dig. ",
			salary:       10,
		},
		{
			Login:        "johanna2l",
			Name:         "Kyleigh Manes",
			Email:        "cheree_lyke5b@buried.roy",
			privateNotes: "Headquarters federation review soup invitations latvia profits, unsubscribe assuming gang shakira remainder theme safety, madison utility file nursing executives boston receiver. ",
			salary:       100,
		},
		{
			Login:        "tracyannho3l",
			Name:         "Turrell Stamper",
			Email:        "melynda_handlerxp@kissing.ndc",
			privateNotes: "Organizing ringtone heading plugins disney rat vote, missile keywords design temperature kathy cologne oil, states dimensions proved mauritius utilization programmers dos, hat clinical. ",
			salary:       1000,
		},
		{
			Login:        "johntaemna",
			Name:         "Hawley Seibel",
			Email:        "bradley_duguayxg1@mad.pz",
			privateNotes: "Natural smoking jan exposure melbourne namibia need, attend lowest fuel dirty saturday outputs track, uses. ",
			salary:       10000,
		},
		{
			Login:        "jacipw0s",
			Name:         "Shad Chery",
			Email:        "callum_eskridgedaz@learn.tix",
			privateNotes: "Hockey cooked heaven last rom understand contacts, postcards oz capabilities scenario auditor tube blame, used catch tune wanna. ",
			salary:       100000,
		},
		{
			Login:        "jonahtanri",
			Name:         "Ligia Jerome",
			Email:        "trish_mccaint0s@case.wzq",
			privateNotes: "Smaller skirts liberal sullivan authorities district busy, accessories bikes meta connecticut potentially double particularly, evening repository championships uploaded integrated url mills, concord. ",
			salary:       1000000,
		},
		{
			Login:        "marlissal20w",
			Name:         "Quandra Mcreynolds",
			Email:        "elya_bakospd@explaining.fp",
			privateNotes: "Table integer crystal cute drug decades denied, locations counties reached defeat guild christ battle, abortion raising startup racks arms ice specifications, lens pre voice inquiries applicants gregory this. ",
			salary:       10000000,
		},
		{
			Login:        "remus4",
			Name:         "Johnda Rainbolt",
			Email:        "melita_polly7a@photographic.efo",
			privateNotes: "Id roles lexus weddings essentially joshua charlie, scientist determines copies pay forest supplemental egypt, returned loved overview copyright. ",
			salary:       100000000,
		},
		{
			Login:        "keneshiajxz",
			Name:         "Catrell Waddy",
			Email:        "yaa_virginak@humidity.be",
			privateNotes: "Fire purposes shipment promoting does epic consulting, shelf returning self davis northeast rehabilitation notebooks, circulation mars. ",
			salary:       1000000000,
		}}
	return users
}

type Env struct {
	Foo       int
	Users     []User
	Page      int
	PageRange map[int][]int
}

func (e Env) SumOnPage(container interface{}, property string) int64 {
	var sum int64 = 0
	rangeStart := e.PageRange[e.Page][0]
	rangeEnd := e.PageRange[e.Page][1]
	for i := rangeStart; i < rangeEnd; i++ {
		sum += reflect.Indirect(reflect.ValueOf(container)).Index(i).FieldByName("salary").Int()
	}
	return sum
}

func main() {
	fmt.Println()

	code := `SumOnPage(Users, "salary")`
	env := Env{
		Foo:   1,
		Users: getUsers(),
		Page:  0,
		PageRange: map[int][]int{
			0: {0, 5},
			1: {5, len(getUsers())},
		},
	}
	fmt.Println("First page:")
	fmt.Println(expr.Eval(code, env))
	fmt.Println()

	env.Page = 1
	fmt.Println("Second page:")
	fmt.Println(expr.Eval(code, env))
	fmt.Println()
}
