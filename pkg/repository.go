package pkg

type Repository interface {
	List() ([]*Package, error)
	Add(pkg *Package) error
	Remove(pkg *Package) error
}
