# gocan

A very streamlined permission library for Golang.

Currently in development, use at your own risk. Any contribution is welcome!

## Install
```
go get github.com/c0m3tx/gocan
```

## Usage
Define a method `Abilities` on your model, which returns a `gocan.Ability` object

```go
type User struct {
}

func (u User) Abilities() gocan.Ability {
  // By default, no ability is set
  var abilities gocan.Ability

  // first parameter is a string, it is possible
  // to use pre-defined gocan constants for basic CRUD functions
  // second parameter is an interface{}, it's possible to use strings
  // or any other object type
  abilities.Grant(gocan.Read, "target", nil)

  // third parameter is an optional comparison function between user and target
  abilities.Grant("update", User{}, reflect.DeepEquals)

  // gocan.Manage is special: includes gocan.Read, gocan.Create,
  // gocan.Update and gocan.Destroy
  abilities.Grant(gocan.Manage, "some other object", nil))

  return abilities
}

func DoSmth() {
  u := User{}
  gocan.Can(u, "read", "target") // => true
  gocan.Can(u, "update", "target") // => false
  gocan.Can(u, "update", u) // => true
  gocan.Can(u, "read", "some other object") // => true
}
```

See `examples` folder for a more detailed example.
