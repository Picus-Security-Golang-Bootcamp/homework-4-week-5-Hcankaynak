package author

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
)

type AuthorRepository struct {
	db *gorm.DB
}

// NewAuthorRepository creating AuthorRepository, migrating it and added data.
func NewAuthorRepository(db *gorm.DB) (*AuthorRepository, error) {
	authorRepo := AuthorRepository{db: db}
	err := authorRepo.db.AutoMigrate(&Author{})
	if err != nil {
		return nil, fmt.Errorf("cannot migrate book repository %v", err)
	}

	sampleBookData, err := getSampleAuthorData()
	if err != nil {
		return nil, fmt.Errorf("cannot init Book Repository %v", err)
	}

	for _, author := range sampleBookData {
		authorRepo.db.Where(Author{Id: author.Id}).Attrs(Author{Id: author.Id, Name: author.Name}).
			FirstOrCreate(&author)
	}
	return &authorRepo, nil
}

// FindAll finding all authors
func (a *AuthorRepository) FindAll() []Author {
	var authors []Author
	a.db.Find(&authors)

	return authors
}

// FindById finding author by id
func (a *AuthorRepository) FindById(id int) *Author {
	var author Author
	a.db.Where(&Author{Id: id}).First(&author)

	return &author
}

// FindByName finding author by name of it.
func (a *AuthorRepository) FindByName(name string) *Author {
	var author Author
	a.db.Where(&Author{Name: name}).First(&author)

	return &author
}

// FindNameByLike finding author by name with Like query.
func (a *AuthorRepository) FindNameByLike(name string) []Author {
	var authors []Author
	a.db.Where("name LIKE ? ", "%"+name+"%").Find(&authors)

	return authors
}

// GetSampleAuthorData reading author json mapping struct and return author list.
func getSampleAuthorData() ([]Author, error) {
	var initialAuthors []Author

	contents, err := ioutil.ReadFile("./data/author.json")
	if err != nil {
		return nil, fmt.Errorf("cannot read 'author.json' %v", err)
	}
	if err := json.Unmarshal(contents, &initialAuthors); err != nil {
		return nil, fmt.Errorf("cannot unmarshall 'author.json' %v", err)
	}
	return initialAuthors, nil
}
