package main

import (
	"fmt"
	"github.com/antonmedv/expr"
	"strconv"
	"strings"
)

type User struct {
	Login        string
	Name         string
	Email        string
	privateNotes string
}

func getUsers() []User {

	users := []User{{
		Login:        "katelynndel",
		Name:         "Tirzah Menges",
		Email:        "elicia_lewteree1@alternative.fbu",
		privateNotes: "Township ping marvel chassis date sorry screensaver, paintings nelson letting lip employ enemy tasks, continues viewing shapes warranty. ",
	},
		{
			Login:        "shellycjpk",
			Name:         "Aryana Renn",
			Email:        "crosby_swiftb@casting.gd",
			privateNotes: "Controversy educational workshop website waters habitat decline, nightmare argue departure eye shapes humidity chelsea, settled dig. ",
		},
		{
			Login:        "johanna2l",
			Name:         "Kyleigh Manes",
			Email:        "cheree_lyke5b@buried.roy",
			privateNotes: "Headquarters federation review soup invitations latvia profits, unsubscribe assuming gang shakira remainder theme safety, madison utility file nursing executives boston receiver. ",
		},
		{
			Login:        "tracyannho3l",
			Name:         "Turrell Stamper",
			Email:        "melynda_handlerxp@kissing.ndc",
			privateNotes: "Organizing ringtone heading plugins disney rat vote, missile keywords design temperature kathy cologne oil, states dimensions proved mauritius utilization programmers dos, hat clinical. ",
		},
		{
			Login:        "johntaemna",
			Name:         "Hawley Seibel",
			Email:        "bradley_duguayxg1@mad.pz",
			privateNotes: "Natural smoking jan exposure melbourne namibia need, attend lowest fuel dirty saturday outputs track, uses. ",
		},
		{
			Login:        "jacipw0s",
			Name:         "Shad Chery",
			Email:        "callum_eskridgedaz@learn.tix",
			privateNotes: "Hockey cooked heaven last rom understand contacts, postcards oz capabilities scenario auditor tube blame, used catch tune wanna. ",
		},
		{
			Login:        "jonahtanri",
			Name:         "Ligia Jerome",
			Email:        "trish_mccaint0s@case.wzq",
			privateNotes: "Smaller skirts liberal sullivan authorities district busy, accessories bikes meta connecticut potentially double particularly, evening repository championships uploaded integrated url mills, concord. ",
		},
		{
			Login:        "marlissal20w",
			Name:         "Quandra Mcreynolds",
			Email:        "elya_bakospd@explaining.fp",
			privateNotes: "Table integer crystal cute drug decades denied, locations counties reached defeat guild christ battle, abortion raising startup racks arms ice specifications, lens pre voice inquiries applicants gregory this. ",
		},
		{
			Login:        "remus4",
			Name:         "Johnda Rainbolt",
			Email:        "melita_polly7a@photographic.efo",
			privateNotes: "Id roles lexus weddings essentially joshua charlie, scientist determines copies pay forest supplemental egypt, returned loved overview copyright. ",
		},
		{
			Login:        "keneshiajxz",
			Name:         "Catrell Waddy",
			Email:        "yaa_virginak@humidity.be",
			privateNotes: "Fire purposes shipment promoting does epic consulting, shelf returning self davis northeast rehabilitation notebooks, circulation mars. ",
		}}
	return users
}

func getUserName(u User) string {
	return u.Name
}

func sumSlice(items []any) any {
	if len(items) == 0 {
		return 0
	}
	var ints int64
	var floats float64
	var hadFloats bool = false
	for _, item := range items {
		switch x := item.(type) {
		case int:
			ints += int64(x)
		case int8:
			ints += int64(x)
		case int16:
			ints += int64(x)
		case int32:
			ints += int64(x)
		case int64:
			ints += x
		case float32:
			hadFloats = true
			floats += float64(x)
		case float64:
			hadFloats = true
			floats += x
		case string:
			if strings.Contains(x, ".") {
				f, err := strconv.ParseFloat(x, 64)
				if err == nil {
					floats += f
				} else {
					panic(err)
				}
			} else {
				i, err := strconv.ParseInt(x, 10, 64)
				if err == nil {
					ints += i
				} else {
					panic(err)
				}
			}
		default:
			panic(fmt.Errorf("cannot call sum() with item of type %T", x))
		}
	}
	if hadFloats {
		return floats + float64(ints)
	}
	return ints
}

func main() {
	env := map[string]interface{}{
		"foo":      1,
		"double":   func(i int) int { return i * 2 },
		"users":    getUsers(),
		"getUsers": getUsers,
		"getName":  getUserName,
		"sum":      sumSlice,
	}

	out, _ := expr.Eval("double(foo)", env)
	fmt.Println(out)
	out, _ = expr.Eval("users[1]", env)
	fmt.Println(out)
	out, _ = expr.Eval("users[0].Name", env)
	fmt.Println(out)
	out, _ = expr.Eval("getUsers()[0].Email", env)
	fmt.Println(out)
	out, _ = expr.Eval("getName(users[2])", env)
	fmt.Println(out)
	out, _ = expr.Eval("1 + 2 + 3.1 + 4", env)
	fmt.Println(out)
	out, _ = expr.Eval("sum([1, 2, 3.1, \"4\"])", env)
	fmt.Println(out)
}
