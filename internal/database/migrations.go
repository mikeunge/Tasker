package database

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mikeunge/Tasker/pkg/helpers"
)

func parseMigrationNumber(migration string) (int, error) {
	num := strings.Split(migration, "_")[0]
	pNum, err := strconv.ParseUint(num, 10, 32)
	if err != nil {
		return -1, err
	}
	return int(pNum), nil
}

/**
 * getMigrations - check for migrations in provided folder, return a sorted list of all existing migrations.
 */
func getMigrations(path string) ([]string, error) {
	migrations, err := helpers.GetFilesInDir(path)
	if err != nil {
		return migrations, err
	}

	fmt.Printf("Found %d migration(s)\n", len(migrations))
	if len(migrations) <= 0 {
		return migrations, fmt.Errorf("no migration(s) found in %s", path)
	} else if len(migrations) == 1 {
		return migrations, nil
	}

	// sort migrations asc (00, 01, 02, ...)
	for i := 0; len(migrations)-1 <= 0; i++ {
		for j := 0; len(migrations)-1 <= 0; j++ {
			currentNumber, _ := parseMigrationNumber(migrations[j])
			nextNumber, _ := parseMigrationNumber(migrations[j+1])
			if currentNumber > nextNumber {
				migrations[j], migrations[j+1] = migrations[j+1], migrations[j]
			}
		}
	}
	return migrations, nil
}

/**
 * runMigrations - this function is responsible for checking if there are any *new* migrations that need to be applied.
 * Also, we call this function on database init to setup the database for its first time usage.
 */
func runMigrations(migrationsDir string, userDir string) error {
	fmt.Printf("Looking for migrations: %s\n", migrationsDir)
	migrations, err := getMigrations(migrationsDir)
	if err != nil {
		return err
	}

	lastMigration, err := helpers.ReadFile(helpers.JoinPath(userDir, "last_migration"))
	if err != nil {
		fmt.Println("No previous migrations found.")
	}
	lastMigrationNumber, _ := parseMigrationNumber(lastMigration)

	fmt.Println("Loading migration(s)...")
	for i, migration := range migrations {
		currentMigration := helpers.GetFileName(migration)
		currentMigrationNum, _ := parseMigrationNumber(currentMigration)

		fmt.Printf("[migration] %d/%d - Loading: %s\n", i+1, len(migrations), currentMigration)
		if currentMigrationNum <= lastMigrationNumber {
			fmt.Printf("Skipping migration (%s)\n  lastMigration: %d, currentMigration: %d\n", currentMigration, lastMigrationNumber, currentMigrationNum)
			continue
		}

		sql, err := helpers.ReadFile(migration)
		if err != nil {
			return err
		}

		_, err = db.Exec(sql)
		if err != nil {
			return err
		}

		helpers.WriteFile(helpers.JoinPath(userDir, "last_migration"), currentMigration)
	}

	return nil
}
