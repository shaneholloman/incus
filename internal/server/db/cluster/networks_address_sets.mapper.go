//go:build linux && cgo && !agent

// Code generated by generate-database from the incus project - DO NOT EDIT.

package cluster

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/mattn/go-sqlite3"
)

var networkAddressSetObjects = RegisterStmt(`
SELECT networks_address_sets.id, networks_address_sets.project_id, projects.name AS project, networks_address_sets.name, coalesce(networks_address_sets.description, ''), networks_address_sets.addresses
  FROM networks_address_sets
  JOIN projects ON networks_address_sets.project_id = projects.id
  ORDER BY projects.id, networks_address_sets.name
`)

var networkAddressSetObjectsByID = RegisterStmt(`
SELECT networks_address_sets.id, networks_address_sets.project_id, projects.name AS project, networks_address_sets.name, coalesce(networks_address_sets.description, ''), networks_address_sets.addresses
  FROM networks_address_sets
  JOIN projects ON networks_address_sets.project_id = projects.id
  WHERE ( networks_address_sets.id = ? )
  ORDER BY projects.id, networks_address_sets.name
`)

var networkAddressSetObjectsByName = RegisterStmt(`
SELECT networks_address_sets.id, networks_address_sets.project_id, projects.name AS project, networks_address_sets.name, coalesce(networks_address_sets.description, ''), networks_address_sets.addresses
  FROM networks_address_sets
  JOIN projects ON networks_address_sets.project_id = projects.id
  WHERE ( networks_address_sets.name = ? )
  ORDER BY projects.id, networks_address_sets.name
`)

var networkAddressSetObjectsByProject = RegisterStmt(`
SELECT networks_address_sets.id, networks_address_sets.project_id, projects.name AS project, networks_address_sets.name, coalesce(networks_address_sets.description, ''), networks_address_sets.addresses
  FROM networks_address_sets
  JOIN projects ON networks_address_sets.project_id = projects.id
  WHERE ( project = ? )
  ORDER BY projects.id, networks_address_sets.name
`)

var networkAddressSetObjectsByProjectAndName = RegisterStmt(`
SELECT networks_address_sets.id, networks_address_sets.project_id, projects.name AS project, networks_address_sets.name, coalesce(networks_address_sets.description, ''), networks_address_sets.addresses
  FROM networks_address_sets
  JOIN projects ON networks_address_sets.project_id = projects.id
  WHERE ( project = ? AND networks_address_sets.name = ? )
  ORDER BY projects.id, networks_address_sets.name
`)

var networkAddressSetID = RegisterStmt(`
SELECT networks_address_sets.id FROM networks_address_sets
  JOIN projects ON networks_address_sets.project_id = projects.id
  WHERE projects.name = ? AND networks_address_sets.name = ?
`)

var networkAddressSetCreate = RegisterStmt(`
INSERT INTO networks_address_sets (project_id, name, description, addresses)
  VALUES ((SELECT projects.id FROM projects WHERE projects.name = ?), ?, ?, ?)
`)

var networkAddressSetRename = RegisterStmt(`
UPDATE networks_address_sets SET name = ? WHERE project_id = (SELECT projects.id FROM projects WHERE projects.name = ?) AND name = ?
`)

var networkAddressSetUpdate = RegisterStmt(`
UPDATE networks_address_sets
  SET project_id = (SELECT projects.id FROM projects WHERE projects.name = ?), name = ?, description = ?, addresses = ?
 WHERE id = ?
`)

var networkAddressSetDeleteByProjectAndName = RegisterStmt(`
DELETE FROM networks_address_sets WHERE project_id = (SELECT projects.id FROM projects WHERE projects.name = ?) AND name = ?
`)

// GetNetworkAddressSetID return the ID of the network_address_set with the given key.
// generator: network_address_set ID
func GetNetworkAddressSetID(ctx context.Context, db tx, project string, name string) (_ int64, _err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	stmt, err := Stmt(db, networkAddressSetID)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"networkAddressSetID\" prepared statement: %w", err)
	}

	row := stmt.QueryRowContext(ctx, project, name)
	var id int64
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return -1, ErrNotFound
	}

	if err != nil {
		return -1, fmt.Errorf("Failed to get \"networks_address_sets\" ID: %w", err)
	}

	return id, nil
}

// NetworkAddressSetExists checks if a network_address_set with the given key exists.
// generator: network_address_set Exists
func NetworkAddressSetExists(ctx context.Context, db dbtx, project string, name string) (_ bool, _err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	stmt, err := Stmt(db, networkAddressSetID)
	if err != nil {
		return false, fmt.Errorf("Failed to get \"networkAddressSetID\" prepared statement: %w", err)
	}

	row := stmt.QueryRowContext(ctx, project, name)
	var id int64
	err = row.Scan(&id)
	if errors.Is(err, sql.ErrNoRows) {
		return false, nil
	}

	if err != nil {
		return false, fmt.Errorf("Failed to get \"networks_address_sets\" ID: %w", err)
	}

	return true, nil
}

// networkAddressSetColumns returns a string of column names to be used with a SELECT statement for the entity.
// Use this function when building statements to retrieve database entries matching the NetworkAddressSet entity.
func networkAddressSetColumns() string {
	return "networks_address_sets.id, networks_address_sets.project_id, projects.name AS project, networks_address_sets.name, coalesce(networks_address_sets.description, ''), networks_address_sets.addresses"
}

// getNetworkAddressSets can be used to run handwritten sql.Stmts to return a slice of objects.
func getNetworkAddressSets(ctx context.Context, stmt *sql.Stmt, args ...any) ([]NetworkAddressSet, error) {
	objects := make([]NetworkAddressSet, 0)

	dest := func(scan func(dest ...any) error) error {
		n := NetworkAddressSet{}
		var addressesStr string
		err := scan(&n.ID, &n.ProjectID, &n.Project, &n.Name, &n.Description, &addressesStr)
		if err != nil {
			return err
		}

		err = unmarshalJSON(addressesStr, &n.Addresses)
		if err != nil {
			return err
		}

		objects = append(objects, n)

		return nil
	}

	err := selectObjects(ctx, stmt, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"networks_address_sets\" table: %w", err)
	}

	return objects, nil
}

// getNetworkAddressSetsRaw can be used to run handwritten query strings to return a slice of objects.
func getNetworkAddressSetsRaw(ctx context.Context, db dbtx, sql string, args ...any) ([]NetworkAddressSet, error) {
	objects := make([]NetworkAddressSet, 0)

	dest := func(scan func(dest ...any) error) error {
		n := NetworkAddressSet{}
		var addressesStr string
		err := scan(&n.ID, &n.ProjectID, &n.Project, &n.Name, &n.Description, &addressesStr)
		if err != nil {
			return err
		}

		err = unmarshalJSON(addressesStr, &n.Addresses)
		if err != nil {
			return err
		}

		objects = append(objects, n)

		return nil
	}

	err := scan(ctx, db, sql, dest, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"networks_address_sets\" table: %w", err)
	}

	return objects, nil
}

// GetNetworkAddressSets returns all available network_address_sets.
// generator: network_address_set GetMany
func GetNetworkAddressSets(ctx context.Context, db dbtx, filters ...NetworkAddressSetFilter) (_ []NetworkAddressSet, _err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	var err error

	// Result slice.
	objects := make([]NetworkAddressSet, 0)

	// Pick the prepared statement and arguments to use based on active criteria.
	var sqlStmt *sql.Stmt
	args := []any{}
	queryParts := [2]string{}

	if len(filters) == 0 {
		sqlStmt, err = Stmt(db, networkAddressSetObjects)
		if err != nil {
			return nil, fmt.Errorf("Failed to get \"networkAddressSetObjects\" prepared statement: %w", err)
		}
	}

	for i, filter := range filters {
		if filter.Project != nil && filter.Name != nil && filter.ID == nil {
			args = append(args, []any{filter.Project, filter.Name}...)
			if len(filters) == 1 {
				sqlStmt, err = Stmt(db, networkAddressSetObjectsByProjectAndName)
				if err != nil {
					return nil, fmt.Errorf("Failed to get \"networkAddressSetObjectsByProjectAndName\" prepared statement: %w", err)
				}

				break
			}

			query, err := StmtString(networkAddressSetObjectsByProjectAndName)
			if err != nil {
				return nil, fmt.Errorf("Failed to get \"networkAddressSetObjects\" prepared statement: %w", err)
			}

			parts := strings.SplitN(query, "ORDER BY", 2)
			if i == 0 {
				copy(queryParts[:], parts)
				continue
			}

			_, where, _ := strings.Cut(parts[0], "WHERE")
			queryParts[0] += "OR" + where
		} else if filter.Project != nil && filter.ID == nil && filter.Name == nil {
			args = append(args, []any{filter.Project}...)
			if len(filters) == 1 {
				sqlStmt, err = Stmt(db, networkAddressSetObjectsByProject)
				if err != nil {
					return nil, fmt.Errorf("Failed to get \"networkAddressSetObjectsByProject\" prepared statement: %w", err)
				}

				break
			}

			query, err := StmtString(networkAddressSetObjectsByProject)
			if err != nil {
				return nil, fmt.Errorf("Failed to get \"networkAddressSetObjects\" prepared statement: %w", err)
			}

			parts := strings.SplitN(query, "ORDER BY", 2)
			if i == 0 {
				copy(queryParts[:], parts)
				continue
			}

			_, where, _ := strings.Cut(parts[0], "WHERE")
			queryParts[0] += "OR" + where
		} else if filter.Name != nil && filter.ID == nil && filter.Project == nil {
			args = append(args, []any{filter.Name}...)
			if len(filters) == 1 {
				sqlStmt, err = Stmt(db, networkAddressSetObjectsByName)
				if err != nil {
					return nil, fmt.Errorf("Failed to get \"networkAddressSetObjectsByName\" prepared statement: %w", err)
				}

				break
			}

			query, err := StmtString(networkAddressSetObjectsByName)
			if err != nil {
				return nil, fmt.Errorf("Failed to get \"networkAddressSetObjects\" prepared statement: %w", err)
			}

			parts := strings.SplitN(query, "ORDER BY", 2)
			if i == 0 {
				copy(queryParts[:], parts)
				continue
			}

			_, where, _ := strings.Cut(parts[0], "WHERE")
			queryParts[0] += "OR" + where
		} else if filter.ID != nil && filter.Name == nil && filter.Project == nil {
			args = append(args, []any{filter.ID}...)
			if len(filters) == 1 {
				sqlStmt, err = Stmt(db, networkAddressSetObjectsByID)
				if err != nil {
					return nil, fmt.Errorf("Failed to get \"networkAddressSetObjectsByID\" prepared statement: %w", err)
				}

				break
			}

			query, err := StmtString(networkAddressSetObjectsByID)
			if err != nil {
				return nil, fmt.Errorf("Failed to get \"networkAddressSetObjects\" prepared statement: %w", err)
			}

			parts := strings.SplitN(query, "ORDER BY", 2)
			if i == 0 {
				copy(queryParts[:], parts)
				continue
			}

			_, where, _ := strings.Cut(parts[0], "WHERE")
			queryParts[0] += "OR" + where
		} else if filter.ID == nil && filter.Name == nil && filter.Project == nil {
			return nil, fmt.Errorf("Cannot filter on empty NetworkAddressSetFilter")
		} else {
			return nil, errors.New("No statement exists for the given Filter")
		}
	}

	// Select.
	if sqlStmt != nil {
		objects, err = getNetworkAddressSets(ctx, sqlStmt, args...)
	} else {
		queryStr := strings.Join(queryParts[:], "ORDER BY")
		objects, err = getNetworkAddressSetsRaw(ctx, db, queryStr, args...)
	}

	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"networks_address_sets\" table: %w", err)
	}

	return objects, nil
}

// GetNetworkAddressSetConfig returns all available NetworkAddressSet Config
// generator: network_address_set GetMany
func GetNetworkAddressSetConfig(ctx context.Context, db tx, networkAddressSetID int, filters ...ConfigFilter) (_ map[string]string, _err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	networkAddressSetConfig, err := GetConfig(ctx, db, "networks_address_sets", "network_address_set", filters...)
	if err != nil {
		return nil, err
	}

	config, ok := networkAddressSetConfig[networkAddressSetID]
	if !ok {
		config = map[string]string{}
	}

	return config, nil
}

// GetNetworkAddressSet returns the network_address_set with the given key.
// generator: network_address_set GetOne
func GetNetworkAddressSet(ctx context.Context, db dbtx, project string, name string) (_ *NetworkAddressSet, _err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	filter := NetworkAddressSetFilter{}
	filter.Project = &project
	filter.Name = &name

	objects, err := GetNetworkAddressSets(ctx, db, filter)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch from \"networks_address_sets\" table: %w", err)
	}

	switch len(objects) {
	case 0:
		return nil, ErrNotFound
	case 1:
		return &objects[0], nil
	default:
		return nil, fmt.Errorf("More than one \"networks_address_sets\" entry matches")
	}
}

// CreateNetworkAddressSet adds a new network_address_set to the database.
// generator: network_address_set Create
func CreateNetworkAddressSet(ctx context.Context, db dbtx, object NetworkAddressSet) (_ int64, _err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	args := make([]any, 4)

	// Populate the statement arguments.
	args[0] = object.Project
	args[1] = object.Name
	args[2] = object.Description
	marshaledAddresses, err := marshalJSON(object.Addresses)
	if err != nil {
		return -1, err
	}

	args[3] = marshaledAddresses

	// Prepared statement to use.
	stmt, err := Stmt(db, networkAddressSetCreate)
	if err != nil {
		return -1, fmt.Errorf("Failed to get \"networkAddressSetCreate\" prepared statement: %w", err)
	}

	// Execute the statement.
	result, err := stmt.Exec(args...)
	var sqliteErr sqlite3.Error
	if errors.As(err, &sqliteErr) {
		if sqliteErr.Code == sqlite3.ErrConstraint {
			return -1, ErrConflict
		}
	}

	if err != nil {
		return -1, fmt.Errorf("Failed to create \"networks_address_sets\" entry: %w", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return -1, fmt.Errorf("Failed to fetch \"networks_address_sets\" entry ID: %w", err)
	}

	return id, nil
}

// CreateNetworkAddressSetConfig adds new network_address_set Config to the database.
// generator: network_address_set Create
func CreateNetworkAddressSetConfig(ctx context.Context, db dbtx, networkAddressSetID int64, config map[string]string) (_err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	referenceID := int(networkAddressSetID)
	for key, value := range config {
		insert := Config{
			ReferenceID: referenceID,
			Key:         key,
			Value:       value,
		}

		err := CreateConfig(ctx, db, "networks_address_sets", "network_address_set", insert)
		if err != nil {
			return fmt.Errorf("Insert Config failed for NetworkAddressSet: %w", err)
		}

	}

	return nil
}

// RenameNetworkAddressSet renames the network_address_set matching the given key parameters.
// generator: network_address_set Rename
func RenameNetworkAddressSet(ctx context.Context, db dbtx, project string, name string, to string) (_err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	stmt, err := Stmt(db, networkAddressSetRename)
	if err != nil {
		return fmt.Errorf("Failed to get \"networkAddressSetRename\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(to, project, name)
	if err != nil {
		return fmt.Errorf("Rename NetworkAddressSet failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows failed: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query affected %d rows instead of 1", n)
	}

	return nil
}

// UpdateNetworkAddressSet updates the network_address_set matching the given key parameters.
// generator: network_address_set Update
func UpdateNetworkAddressSet(ctx context.Context, db tx, project string, name string, object NetworkAddressSet) (_err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	id, err := GetNetworkAddressSetID(ctx, db, project, name)
	if err != nil {
		return err
	}

	stmt, err := Stmt(db, networkAddressSetUpdate)
	if err != nil {
		return fmt.Errorf("Failed to get \"networkAddressSetUpdate\" prepared statement: %w", err)
	}

	marshaledAddresses, err := marshalJSON(object.Addresses)
	if err != nil {
		return err
	}

	result, err := stmt.Exec(object.Project, object.Name, object.Description, marshaledAddresses, id)
	if err != nil {
		return fmt.Errorf("Update \"networks_address_sets\" entry failed: %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n != 1 {
		return fmt.Errorf("Query updated %d rows instead of 1", n)
	}

	return nil
}

// UpdateNetworkAddressSetConfig updates the network_address_set Config matching the given key parameters.
// generator: network_address_set Update
func UpdateNetworkAddressSetConfig(ctx context.Context, db tx, networkAddressSetID int64, config map[string]string) (_err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	err := UpdateConfig(ctx, db, "networks_address_sets", "network_address_set", int(networkAddressSetID), config)
	if err != nil {
		return fmt.Errorf("Replace Config for NetworkAddressSet failed: %w", err)
	}

	return nil
}

// DeleteNetworkAddressSet deletes the network_address_set matching the given key parameters.
// generator: network_address_set DeleteOne-by-Project-and-Name
func DeleteNetworkAddressSet(ctx context.Context, db dbtx, project string, name string) (_err error) {
	defer func() {
		_err = mapErr(_err, "Network_address_set")
	}()

	stmt, err := Stmt(db, networkAddressSetDeleteByProjectAndName)
	if err != nil {
		return fmt.Errorf("Failed to get \"networkAddressSetDeleteByProjectAndName\" prepared statement: %w", err)
	}

	result, err := stmt.Exec(project, name)
	if err != nil {
		return fmt.Errorf("Delete \"networks_address_sets\": %w", err)
	}

	n, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("Fetch affected rows: %w", err)
	}

	if n == 0 {
		return ErrNotFound
	} else if n > 1 {
		return fmt.Errorf("Query deleted %d NetworkAddressSet rows instead of 1", n)
	}

	return nil
}
