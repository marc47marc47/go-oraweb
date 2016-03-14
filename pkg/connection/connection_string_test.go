package connection

import (
	"fmt"
	"os/user"
	"testing"

	"github.com/stretchr/testify/assert"
	"pkg/command"
)

func Test_Invalid_Url(t *testing.T) {
	opts := command.Options{}
	examples := []string{
		"postgre://foobar",
		"foobar",
	}

	for _, val := range examples {
		opts.Url = val
		str, err := BuildString(opts)

		assert.Equal(t, "", str)
		assert.Error(t, err)
		assert.Equal(t, "Invalid URL. Valid format: oracle://user:password@host:port/db?sslmode=mode", err.Error())
	}
}

func Test_Valid_Url(t *testing.T) {
	url := "oracle://myhost/database"
	str, err := BuildString(command.Options{Url: url})

	assert.Equal(t, nil, err)
	assert.Equal(t, url, str)
}

func Test_Url_And_Ssl_Flag(t *testing.T) {
	str, err := BuildString(command.Options{
		Url: "oracle://myhost/database",
		Ssl: "disable",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://myhost/database?sslmode=disable", str)
}

func Test_Localhost_Url_And_No_Ssl_Flag(t *testing.T) {
	str, err := BuildString(command.Options{
		Url: "oracle://localhost/database",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://localhost/database?sslmode=disable", str)

	str, err = BuildString(command.Options{
		Url: "oracle://127.0.0.1/database",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://127.0.0.1/database?sslmode=disable", str)
}

func Test_Localhost_Url_And_Ssl_Flag(t *testing.T) {
	str, err := BuildString(command.Options{
		Url: "oracle://localhost/database",
		Ssl: "require",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://localhost/database?sslmode=require", str)

	str, err = BuildString(command.Options{
		Url: "postgres://127.0.0.1/database",
		Ssl: "require",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://127.0.0.1/database?sslmode=require", str)
}

func Test_Localhost_Url_And_Ssl_Arg(t *testing.T) {
	str, err := BuildString(command.Options{
		Url: "oracle://localhost/database?sslmode=require",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://localhost/database?sslmode=require", str)

	str, err = BuildString(command.Options{
		Url: "oracle://127.0.0.1/database?sslmode=require",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://127.0.0.1/database?sslmode=require", str)
}

func Test_Flag_Args(t *testing.T) {
	str, err := BuildString(command.Options{
		Host:   "host",
		Port:   1522,
		User:   "user",
		Pass:   "password",
		DbName: "db",
	})

	assert.Equal(t, nil, err)
	assert.Equal(t, "postgres://user:password@host:5432/db", str)
}

func Test_Localhost(t *testing.T) {
	opts := command.Options{
		Host:   "localhost",
		Port:   1522,
		User:   "user",
		Pass:   "password",
		DbName: "db",
	}

	str, err := BuildString(opts)
	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://user:password@localhost:5432/db?sslmode=disable", str)

	opts.Host = "127.0.0.1"
	str, err = BuildString(opts)
	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://user:password@127.0.0.1:5432/db?sslmode=disable", str)
}

func Test_Localhost_And_Ssl(t *testing.T) {
	opts := command.Options{
		Host:   "localhost",
		Port:   1522,
		User:   "user",
		Pass:   "password",
		DbName: "db",
		Ssl:    "require",
	}

	str, err := BuildString(opts)
	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://user:password@localhost:5432/db?sslmode=require", str)
}

func Test_No_User(t *testing.T) {
	opts := command.Options{Host: "host", Port: 5432, DbName: "db"}
	u, _ := user.Current()
	str, err := BuildString(opts)

	assert.Equal(t, nil, err)
	assert.Equal(t, fmt.Sprintf("oracle://%s@host:5432/db", u.Username), str)
}

func Test_Port(t *testing.T) {
	opts := command.Options{Host: "host", User: "user", Port: 5000, DbName: "db"}
	str, err := BuildString(opts)

	assert.Equal(t, nil, err)
	assert.Equal(t, "oracle://user@host:5000/db", str)
}

func Test_Blank(t *testing.T) {
	assert.Equal(t, true, IsBlank(command.Options{}))
	assert.Equal(t, false, IsBlank(command.Options{Host: "host", User: "user"}))
	assert.Equal(t, false, IsBlank(command.Options{Host: "host", User: "user", DbName: "db"}))
	assert.Equal(t, false, IsBlank(command.Options{Url: "url"}))
}
