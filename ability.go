package gocan

const (
	// Manage includes Read, Create, Update and Destroy
	Manage = "manage"

	// Read is the basic read-only permission in CRUD
	Read = "read"

	// Create is the basic create permission in CRUD
	Create = "create"

	// Update is the basic update permission in CRUD
	Update = "update"

	// Destroy is the basic removal permission in CRUD
	Destroy = "destroy"
)

// Capable represents objects which have a list of abilities
type Capable interface {
	Abilities() Ability
}

// An ability represents the single ability (with a permission/target and optional compare function)
type ability struct {
	Permission string
	Target     interface{}
	Compare    func(interface{}, interface{}) bool
}

// An Ability is a set of single abilities (A)
type Ability []ability

func (a Ability) can(permission string, target interface{}) bool {
	for _, ab := range a {
		if ab.Permission == permission {
			if target == nil || ab.Target == nil {
				return true
			}

			if ab.Compare(ab.Target, target) {
				return true
			}
		}
	}

	return false
}

// Grant adds a new ability with permission, an (optional) target and an (optional) compare function.
//
// If target is not supplied, the grant is applied to all system objects.
//
// If compare function is not supplied, targets are compared with equality, if present.
func (a *Ability) Grant(permission string, target interface{}, compare func(interface{}, interface{}) bool) {
	if compare == nil {
		compare = basicEquality
	}

	if permission == Manage {
		a.Grant(Read, target, compare)
		a.Grant(Create, target, compare)
		a.Grant(Update, target, compare)
		a.Grant(Destroy, target, compare)
		return
	}

	*a = append(*a, ability{Permission: permission, Target: target, Compare: compare})
}

// Deny removes an existing ability.
// Useful in combination with Manage, which adds all the basic CRUD functions.
//
// If compare function is not supplied, targets are compared with equality, if present.
func (a *Ability) Deny(permission string, target interface{}, compare func(interface{}, interface{}) bool) {
	if compare == nil {
		compare = basicEquality
	}

	for i, ab := range *a {
		if ab.Permission == permission && compare(ab.Target, target) {
			*a = append((*a)[0:i], (*a)[i+1:]...)
			return
		}
	}
}

// Can allows verification of a capable user for some permission/target pair
func Can(ab Capable, permission string, target interface{}) bool {
	return ab.Abilities().can(permission, target)
}

func basicEquality(x interface{}, y interface{}) bool {
	return x == y
}
