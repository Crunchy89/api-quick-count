package repository_test

import (
	"testing"

	"github.com/crunchy89/api-quick-count/app/domain/entities"
	"github.com/crunchy89/api-quick-count/app/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRoleRepository(t *testing.T) {
	// Set up an in-memory SQLite database for testing
	db, err := ConnectingToDatabaseTest()
	if err != nil {
		panic(err)
	}
	require.NoError(t, err)

	// AutoMigrate creates the necessary table
	v := new(entities.Role)
	if !db.Migrator().HasTable(v) {
		db.Debug().AutoMigrate(v)
	} else {
		db.Debug().Migrator().DropTable(v)
		db.Debug().AutoMigrate(v)
	}

	// Create a new RoleRepository instance for testing
	roleRepo := repository.NewRoleRepository(db)

	// Test Save function
	t.Run("Save", func(t *testing.T) {
		// Create a role
		role := &entities.Role{
			Name: "Admin",
		}

		// Save the role
		id, err := roleRepo.Save(role)
		require.NoError(t, err)

		// Check that the role was saved with a valid ID
		assert.NotNil(t, id)
		assert.NotZero(t, *id)
	})

	// Test Create function
	t.Run("Create", func(t *testing.T) {
		// Create a role
		role := &entities.Role{Name: "TestRole"}

		// Create the role
		uuid, err := roleRepo.Create(role)
		require.NoError(t, err)

		// Check that the role was created with a valid UUID
		assert.NotNil(t, uuid)
		assert.NotEmpty(t, *uuid)
	})

	// Test GetAll function
	t.Run("GetAll", func(t *testing.T) {

		// Get all roles
		allRoles, err := roleRepo.GetAll()
		require.NoError(t, err)

		// Check that the number of retrieved roles matches the number of saved roles
		assert.Greater(t, len(allRoles), 0)
	})

	// Test GetByID function
	t.Run("GetByID", func(t *testing.T) {
		// Create a role
		role := &entities.Role{Name: "TestRole"}
		_, err := roleRepo.Save(role)
		require.NoError(t, err)

		// Get the role by ID
		retrievedRole, err := roleRepo.GetByID(role.ID)
		require.NoError(t, err)

		// Check that the retrieved role matches the saved role
		assert.Equal(t, role.Name, retrievedRole.Name)
	})

	// Test GetByUUID function
	t.Run("GetByUUID", func(t *testing.T) {
		// Create a role
		role := &entities.Role{Name: "TestRole"}
		_, err := roleRepo.Save(role)
		require.NoError(t, err)

		// Get the role by UUID
		retrievedRole, err := roleRepo.GetByUUID(role.UUID)
		require.NoError(t, err)

		// Check that the retrieved role matches the saved role
		assert.Equal(t, role.Name, retrievedRole.Name)
	})
	t.Run("UpdateByUuid", func(t *testing.T) {
		// Create a role
		role := &entities.Role{Name: "TestRoleID"}

		// Create the role
		_, err := roleRepo.Save(role)
		require.NoError(t, err)

		role.Name = "TestUpdateRoleID"
		err = roleRepo.UpdateById(role)
		require.NoError(t, err)
		assert.NoError(t, err)

		result, err := roleRepo.GetByID(role.ID)
		require.NoError(t, err)
		assert.Equal(t, role.Name, result.Name)
	})
	t.Run("UpdateByUuid", func(t *testing.T) {
		// Create a role
		role := &entities.Role{Name: "TestRoleUUID"}

		// Create the role
		_, err := roleRepo.Create(role)
		require.NoError(t, err)

		role.Name = "TestUpdateRoleUUID"
		err = roleRepo.UpdateByUuid(role)
		require.NoError(t, err)
		assert.NoError(t, err)

		result, err := roleRepo.GetByUUID(role.UUID)
		require.NoError(t, err)
		assert.Equal(t, role.Name, result.Name)
	})
	t.Run("DeleteById", func(t *testing.T) {
		// Create a role
		role := &entities.Role{Name: "TestRoleDeleteByID"}

		// Create the role
		_, err := roleRepo.Save(role)
		require.NoError(t, err)
		err = roleRepo.DeleteById(role.ID)
		require.NoError(t, err)
		assert.NoError(t, err)

		_, err = roleRepo.GetByID(role.ID)
		assert.Error(t, err)
	})
	t.Run("DeleteByUuid", func(t *testing.T) {
		// Create a role
		role := &entities.Role{Name: "TestRoleDeleteByUUID"}

		// Create the role
		uuid, err := roleRepo.Create(role)
		require.NoError(t, err)
		err = roleRepo.DeleteByUuid(*uuid)
		require.NoError(t, err)
		assert.NoError(t, err)

		_, err = roleRepo.GetByUUID(*uuid)
		assert.Error(t, err)
	})
	t.Run("Done testing drop table", func(t *testing.T) {
		err := db.Migrator().DropTable(&entities.Role{})
		assert.NoError(t, err)
	})

}
