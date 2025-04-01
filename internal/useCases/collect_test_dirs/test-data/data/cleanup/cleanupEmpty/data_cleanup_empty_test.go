package dataUpCorrect

import (
	"dbevo/commands/commandsData/commandsDataCleanupDeletedScriptsInDb"
	"dbevo/commands/commandsData/commandsDataUp"
	"dbevo/commands/commandsDbEvo"
	"dbevo/commands/commandsMigrations/commandsMigrationsUp"
	"dbevo/config"
	"dbevo/dbTesting"
	"dbevo/dbdrivers/postgresql"
	"dbevo/e2eTesting"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDataUpShouldReturnInsertedDataScriptsAndNoError(t *testing.T) {
	// setup init
	postgresConfig := config.ReadConfig(e2eTesting.ConfigPathFourDirsUp)
	postgresql.OpenPostgreSqlConnection(postgresConfig)
	err := commandsDbEvo.Init()
	assert.Nil(t, err)

	// setup dirs
	migrationsUpDir := "test-data/migrations"
	dataUpDir := "test-data/data"

	// setup migrations
	err = commandsMigrationsUp.Up(migrationsUpDir)
	assert.NoError(t, err)

	// setup data up
	_, err = commandsDataUp.Up(dataUpDir)
	assert.NoError(t, err)

	// when
	deletedDataPaths, err := commandsDataCleanupDeletedScriptsInDb.CleanupDeletedScriptsInDb(dataUpDir)

	// assert err
	assert.NoError(t, err)

	// assert expected
	assert.Empty(t, deletedDataPaths)

	// cleanup
	err = dbTesting.DropTableIfExists("example")
	assert.NoError(t, err)
	err = dbTesting.DropTableIfExists("example3")
	assert.NoError(t, err)
	err = commandsDbEvo.DropDbEvo()
	assert.NoError(t, err)
}
