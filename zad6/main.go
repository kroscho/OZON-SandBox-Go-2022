package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

type User struct {
	ID              int
	Friends         []int
	PossibleFriends []int
}

func NewUser(id int, friends []int, possibleFriends []int) *User {
	return &User{
		ID:              id,
		Friends:         friends,
		PossibleFriends: possibleFriends,
	}
}

func (u *User) addFriend(friendID int) {
	u.Friends = append(u.Friends, friendID)
}

func inputInt(in *bufio.Reader, text string) (int, int) {
	var x, y int
	fmt.Print(text)
	fmt.Fscan(in, &x, &y)

	return x, y
}

func preparePair(usersMap map[int]*User, userID1 int, userID2 int) {
	u := usersMap[userID1]

	u.addFriend(userID2)
	sort.Ints(u.Friends)
	usersMap[userID1] = u
}

func addUsers(usersMap map[int]*User, countUsers int) {
	for i := 0; i < countUsers; i++ {
		id := i + 1
		usersMap[id] = NewUser(id, []int{}, []int{})
	}
}

func (u *User) isFriend(curUser User) bool {
	for _, friendID := range curUser.Friends {
		if u.ID == friendID {
			return true
		}
	}
	return false
}

func (u *User) countGeneralFriends(curUser User) int {
	count := 0
	for i := 0; i < int(math.Min(float64(len(curUser.Friends)), float64(len(u.Friends)))); i++ {
		if u.Friends[i] == curUser.Friends[i] {
			count += 1
		}
	}
	return count
}

func printResult(out *bufio.Writer, users []*User) {
	for _, u := range users {
		pf := u.PossibleFriends

		if len(pf) == 0 {
			fmt.Fprintln(out, 0)
		} else {
			sort.Slice(pf, func(i, j int) (less bool) {
				return pf[i] < pf[j]
			})
			for _, f := range pf {
				fmt.Fprint(out, f, " ")
			}
			fmt.Fprintln(out, "")
		}
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	countUsers, countPairs := inputInt(in, "Введите кол-во пользователей и пар: ")
	usersMap := make(map[int]*User, countUsers)
	users := make([]*User, 0)

	addUsers(usersMap, countUsers)

	for i := 0; i < countPairs; i++ {
		userID1, userID2 := inputInt(in, "Введите пару: ")

		preparePair(usersMap, userID1, userID2)
		preparePair(usersMap, userID2, userID1)
	}

	for _, u := range usersMap {
		users = append(users, u)
	}

	sort.Slice(users, func(i, j int) (less bool) {
		return len(users[i].Friends) > len(users[j].Friends)
	})

	for i, curUser := range users {
		friends := curUser.Friends
		prevCountFriends := len(friends)
		maxGeneralFriends := 0

		for _, u := range users {
			if u.ID == curUser.ID {
				prevCountFriends = len(u.Friends)
				continue
			}
			if u.isFriend(*curUser) {
				prevCountFriends = len(u.Friends)
				continue
			}
			if (len(u.Friends) < prevCountFriends) && (maxGeneralFriends > len(u.Friends)) {
				break
			}
			countGeneralFriends := u.countGeneralFriends(*curUser)
			if countGeneralFriends > maxGeneralFriends {
				maxGeneralFriends = countGeneralFriends
				users[i].PossibleFriends = []int{u.ID}
			} else if countGeneralFriends == maxGeneralFriends && countGeneralFriends != 0 {
				users[i].PossibleFriends = append(curUser.PossibleFriends, u.ID)
			}
			prevCountFriends = len(u.Friends)
		}
	}

	sort.Slice(users, func(i, j int) (less bool) {
		return users[i].ID < users[j].ID
	})

	printResult(out, users)
}
