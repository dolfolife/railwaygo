package collect

import (
	"errors"
	"fmt"
	"github.com/dolfolife/railwaygo/pkg/mapper"
	"github.com/dolfolife/railwaygo/pkg/result"
	"testing"
)

func TestCollectEmpty(t *testing.T) {

	numbers := []int{}
	s := result.NewSliceResult[int](numbers)
	nToCollect := mapper.Map(s, func(r result.Result[int]) int {
		return r.Val
	})
	actualValues := Collect(nToCollect)
	if len(numbers) != len(actualValues) {
		t.Errorf("there was an error with Collect on empty ")
	}
}

func TestCollectSome_MaintainsOrder(t *testing.T) {

	numbers := []int{1, 2, 3, 4, 5}
	s := result.NewSliceResult[int](numbers)
	nToCollect := mapper.Map(s, Fold)
	actualValues := Collect(nToCollect)

	if len(numbers) != len(actualValues) {
		t.Errorf("there was an error with Collect")
	}

	for idx, v := range actualValues {
		if numbers[idx] != v {
			t.Errorf("there was an error with Collect")
		}
	}
}

func TestCollectSome_FlowsEvenWithErrorData(t *testing.T) {

	ids := []int{1, 2, 3, 4, 5}
	users := []User{{Id: 1, Name: "bob"}, {Id: 2, Name: "bob"}, {Id: 3, Name: "bob"}, {Id: 5, Name: "bob"}}
	s := result.NewSliceResult[int](ids)
	usersFound := mapper.Map(s, func(v result.Result[int]) result.Result[User] {
		for _, user := range users {
			if user.Id == v.Val {
				return result.Result[User]{Val: user, Error: v.Error}
			}
		}
		return result.Result[User]{Error: errors.Join(v.Error, errors.New(fmt.Sprintf("user not found with id %d", v.Val)))}
	})
	nToCollect := mapper.Map(usersFound, Fold)
	actualUsers := Collect(nToCollect)

	if len(ids) != len(actualUsers) {
		t.Errorf("there was an error with Collect")
	}

	idWithError := 4
	for _, u := range actualUsers {
		if u.Id == idWithError {
			t.Errorf("found user that it should not have")
		}
	}
}

type User struct {
	Id   int
	Name string
}
