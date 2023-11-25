package godb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type Instance struct {
	Db *pgxpool.Pool
}

func (i *Instance) Start() {
	fmt.Println("Project godb started!")

	//i.addUser(context.Background(), "Walter", 52, true)
	//i.addUser(context.Background(), "Jesse", 21, true)
	//i.getAllUsers(context.Background())
	//i.updateUserAge(context.Background(), "Jesse", 22)
	//i.getAllUsers(context.Background())
	//i.getUserByName(context.Background(), "Walter")

	//i.addRandomUsersFromFile(context.Background(), "module08/internal/godb/users")
	//i.getAllUsers(context.Background())

}

func (i *Instance) addUser(ctx context.Context, name string, age int, isVerify bool) {
	commandTag, err := i.Db.Exec(ctx, "INSERT INTO users (created_at, name, age, verify) VALUES "+
		"($1, $2, $3, $4)", time.Now(), name, age, isVerify)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(commandTag.String())
	fmt.Println(commandTag.RowsAffected())
}

func (i *Instance) getAllUsers(ctx context.Context) {
	var users []User

	rows, err := i.Db.Query(ctx, "SELECT name, age, verify FROM users;")
	if err == pgx.ErrNoRows {
		fmt.Println("No rows!")
		return
	} else if err != nil {
		fmt.Println(err)
		return
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		err = rows.Scan(&user.Name, &user.Age, &user.IsVerify)
		if err != nil {
			fmt.Println(err)
			return
		}
		users = append(users, user)
	}

	fmt.Println(users)
}

func (i *Instance) updateUserAge(ctx context.Context, name string, newAge int) {
	commandTag, err := i.Db.Exec(ctx, "UPDATE users SET age = $2 WHERE name = $1", name, newAge)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(commandTag.String())
}

func (i *Instance) getUserByName(ctx context.Context, name string) {
	row := i.Db.QueryRow(ctx, "SELECT name, age, verify FROM users WHERE name = $1", name)

	user := &User{}
	err := row.Scan(&user.Name, &user.Age, &user.IsVerify)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(user)
}

func (i *Instance) addRandomUsersFromFile(ctx context.Context, filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
		return
	}

	names := strings.Split(string(bytes), "\n")
	for _, name := range names {
		name = strings.Split(name, "\r")[0]
		age := rand.Intn(32) + 18     // random age
		isVerify := rand.Intn(2) == 1 // random boolean

		i.addUser(ctx, name, age, isVerify)
	}
}
