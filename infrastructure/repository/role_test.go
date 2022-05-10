package repository_test

import (
	"errors"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/eduardohslfreire/animalia-api/entity"
	"github.com/eduardohslfreire/animalia-api/infrastructure/repository"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestFindByID(t *testing.T) {
	db, mockSql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to simulate a database connection: %s", err.Error())
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	t.Run("success - find role by id", func(t *testing.T) {
		mockRole := &entity.Role{ID: 1, Name: "Civil", Single: false}
		rows := mockSql.NewRows([]string{"id", "name", "single"}).AddRow(mockRole.ID, mockRole.Name, mockRole.Single)
		query := regexp.QuoteMeta(`SELECT * FROM "roles" WHERE "roles"."id" = $1 ORDER BY "roles"."id" LIMIT 1`)
		mockSql.ExpectQuery(query).WithArgs(1).WillReturnRows(rows)

		roleRepository := repository.NewRoleRepository(gormDB)
		role, err := roleRepository.FindByID(1)
		assert.NotNil(t, role)
		assert.NoError(t, err)
	})

	t.Run("error - find role by id", func(t *testing.T) {
		query := regexp.QuoteMeta(`SELECT * FROM "roles" WHERE "roles"."id" = $1 ORDER BY "roles"."id" LIMIT 1`)
		mockSql.ExpectQuery(query).WithArgs(1).WillReturnError(errors.New("Error"))

		roleRepository := repository.NewRoleRepository(gormDB)
		role, err := roleRepository.FindByID(1)
		assert.Nil(t, role)
		assert.Error(t, err)
		assert.Equal(t, "Error", err.Error())
	})
}

func TestFindAll(t *testing.T) {
	db, mockSql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to simulate a database connection: %s", err.Error())
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	t.Run("success - find all roles", func(t *testing.T) {
		mockRole := entity.Role{ID: 1, Name: "Civil", Single: false}
		rows := mockSql.NewRows([]string{"id", "name", "single"}).AddRow(mockRole.ID, mockRole.Name, mockRole.Single)
		query := regexp.QuoteMeta(`SELECT * FROM "roles"`)
		mockSql.ExpectQuery(query).WithArgs().WillReturnRows(rows)

		roleRepository := repository.NewRoleRepository(gormDB)
		roles, err := roleRepository.FindAll()
		assert.NotNil(t, roles)
		assert.Equal(t, 1, len(*roles))
		assert.NoError(t, err)
	})

	t.Run("error - find all roles", func(t *testing.T) {
		query := regexp.QuoteMeta(`SELECT * FROM "roles"`)
		mockSql.ExpectQuery(query).WithArgs().WillReturnError(errors.New("Error"))

		roleRepository := repository.NewRoleRepository(gormDB)
		role, err := roleRepository.FindAll()
		assert.Nil(t, role)
		assert.Error(t, err)
		assert.Equal(t, "Error", err.Error())
	})
}

func TestFindAllCitizensByID(t *testing.T) {
	db, mockSql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to simulate a database connection: %s", err.Error())
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	t.Run("success - find all citizens by role id", func(t *testing.T) {
		mockCitizen := &entity.Citizen{ID: 1, Name: "Name", Species: "Species", Description: "Description", PhotoURL: "PhotoURL", Weight: 0, Height: 0, HasPetHuman: true}
		rowsCitizen := mockSql.NewRows([]string{"id", "name", "species", "description", "photo_url", "weight", "height", "has_pet_human"}).AddRow(mockCitizen.ID, mockCitizen.Name, mockCitizen.Species, mockCitizen.Description, mockCitizen.PhotoURL, mockCitizen.Weight, mockCitizen.Height, mockCitizen.HasPetHuman)
		query := regexp.QuoteMeta(`SELECT "citizens"."id","citizens"."name","citizens"."species","citizens"."description","citizens"."photo_url","citizens"."weight","citizens"."height","citizens"."has_pet_human" FROM "citizens" JOIN "citizen_role" ON "citizen_role"."citizen_id" = "citizens"."id" AND "citizen_role"."role_id" = $1`)
		mockSql.ExpectQuery(query).WithArgs(1).WillReturnRows(rowsCitizen)

		roleRepository := repository.NewRoleRepository(gormDB)
		citizens, err := roleRepository.FindAllCitizensByID(1)
		assert.NotNil(t, citizens)
		assert.NoError(t, err)
	})

	t.Run("error - find all citizens by role id", func(t *testing.T) {
		query := regexp.QuoteMeta(`SELECT "citizens"."id","citizens"."name","citizens"."species","citizens"."description","citizens"."photo_url","citizens"."weight","citizens"."height","citizens"."has_pet_human" FROM "citizens" JOIN "citizen_role" ON "citizen_role"."citizen_id" = "citizens"."id" AND "citizen_role"."role_id" = $1`)
		mockSql.ExpectQuery(query).WithArgs(1).WillReturnError(errors.New("Error"))

		roleRepository := repository.NewRoleRepository(gormDB)
		role, err := roleRepository.FindAllCitizensByID(1)
		assert.Nil(t, role)
		assert.Error(t, err)
		assert.Equal(t, "Error", err.Error())
	})
}

func TestCountAssociations(t *testing.T) {
	db, mockSql, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to simulate a database connection: %s", err.Error())
	}

	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	t.Run("success - count citizens associated role", func(t *testing.T) {
		row := mockSql.NewRows([]string{"count"}).AddRow(1)
		query := regexp.QuoteMeta(`SELECT count(*) FROM "citizens" JOIN "citizen_role" ON "citizen_role"."citizen_id" = "citizens"."id" AND "citizen_role"."role_id" = $1`)
		mockSql.ExpectQuery(query).WithArgs(1).WillReturnRows(row)

		roleRepository := repository.NewRoleRepository(gormDB)
		countCitizens := roleRepository.CountAssociations(1)
		assert.Equal(t, int64(1), countCitizens)
	})
}
