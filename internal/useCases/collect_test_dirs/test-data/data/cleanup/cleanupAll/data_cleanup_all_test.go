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
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestDataUpShouldReturnInsertedDataScriptsAndNoError(t *testing.T) {
	// setup init
	postgresConfig := config.ReadConfig(e2eTesting.ConfigPathFourDirsUp)
	postgresql.OpenPostgreSqlConnection(postgresConfig)
	err := commandsDbEvo.Init()
	require.NoError(t, err)

	// setup dirs
	migrationsUpDir := "test-data/migrations"
	dataUpDir := "test-data/data"

	// setup file paths
	exampleFilepath := dataUpDir + "/example/D-Insert-Into-example.sql"
	exampleFilepathNew := dataUpDir + "/example/D-Insert-Into-example-new.sql"
	example3Filepath := dataUpDir + "/example3/D-Insert-Into-example3.sql"
	example3FilepathNew := dataUpDir + "/example3/D-Insert-Into-example3-new.sql"

	// setup migrations
	err = commandsMigrationsUp.Up(migrationsUpDir)
	assert.NoError(t, err)

	// setup data up
	_, err = commandsDataUp.Up(dataUpDir)
	assert.NoError(t, err)

	//then change filepaths by renaming the file
	err = os.Rename(exampleFilepath, exampleFilepathNew)
	assert.NoError(t, err)

	err = os.Rename(example3Filepath, example3FilepathNew)
	assert.NoError(t, err)

	// when
	deletedDataPaths, err := commandsDataCleanupDeletedScriptsInDb.CleanupDeletedScriptsInDb(dataUpDir)

	// then assert err
	assert.NoError(t, err)

	// then assert expected
	expected := []string{
		exampleFilepath,
		example3Filepath,
	}
	assert.ElementsMatch(t, deletedDataPaths, expected)

	// then assert data_checksum table is empty
	dataChecksumIsEmpty, err := dbTesting.TableIsEmpty("dbevo.data_checksum")
	assert.NoError(t, err)

	assert.True(t, dataChecksumIsEmpty)

	// reset files
	err = os.Rename(exampleFilepathNew, exampleFilepath)
	assert.NoError(t, err)

	err = os.Rename(example3FilepathNew, example3Filepath)
	assert.NoError(t, err)

	// cleanup db
	err = dbTesting.DropTableIfExists("example")
	assert.NoError(t, err)
	err = dbTesting.DropTableIfExists("example3")
	assert.NoError(t, err)
	err = commandsDbEvo.DropDbEvo()
	assert.NoError(t, err)

}
