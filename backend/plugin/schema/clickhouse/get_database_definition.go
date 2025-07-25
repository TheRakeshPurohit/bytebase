package clickhouse

import (
	"fmt"
	"io"
	"strings"

	storepb "github.com/bytebase/bytebase/backend/generated-go/store"
	"github.com/bytebase/bytebase/backend/plugin/schema"
)

func init() {
	schema.RegisterGetDatabaseDefinition(storepb.Engine_CLICKHOUSE, GetDatabaseDefinition)
	schema.RegisterGetTableDefinition(storepb.Engine_CLICKHOUSE, GetTableDefinition)
	schema.RegisterGetViewDefinition(storepb.Engine_CLICKHOUSE, GetViewDefinition)
}

func GetDatabaseDefinition(_ schema.GetDefinitionContext, to *storepb.DatabaseSchemaMetadata) (string, error) {
	toState := convertToDatabaseState(to)
	var sb strings.Builder

	if err := writeTables(&sb, to, toState); err != nil {
		return "", err
	}
	if err := writeViews(&sb, to, toState); err != nil {
		return "", err
	}

	// Make goyamlv3 happy.
	return strings.TrimLeft(sb.String(), "\n"), nil
}

func GetTableDefinition(_ string, table *storepb.TableMetadata, _ []*storepb.SequenceMetadata) (string, error) {
	var buf strings.Builder
	tableState := convertToTableState(0, table)

	if _, err := buf.WriteString(getTableAnnouncement(table.Name)); err != nil {
		return "", err
	}
	if err := tableState.toString(&buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func GetViewDefinition(_ string, view *storepb.ViewMetadata) (string, error) {
	var buf strings.Builder
	viewState := convertToViewState(0, view)

	if _, err := buf.WriteString(getViewAnnouncement(view.Name)); err != nil {
		return "", err
	}
	if err := viewState.toString(&buf); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func writeTables(w io.StringWriter, to *storepb.DatabaseSchemaMetadata, state *databaseState) error {
	// Follow the order of the input schemas.
	for _, schema := range to.Schemas {
		schemaState, ok := state.schemas[schema.Name]
		if !ok {
			continue
		}
		// Follow the order of the input tables.
		for _, table := range schema.Tables {
			table, ok := schemaState.tables[table.Name]
			if !ok {
				continue
			}
			if _, err := w.WriteString(getTableAnnouncement(table.name)); err != nil {
				return err
			}

			buf := &strings.Builder{}
			if err := table.toString(buf); err != nil {
				return err
			}
			if _, err := w.WriteString(buf.String()); err != nil {
				return err
			}
			delete(schemaState.tables, table.name)
		}
	}
	return nil
}

func writeViews(w io.StringWriter, to *storepb.DatabaseSchemaMetadata, state *databaseState) error {
	// Follow the order of the input schemas.
	for _, schema := range to.Schemas {
		schemaState, ok := state.schemas[schema.Name]
		if !ok {
			continue
		}
		// Follow the order of the input views.
		for _, view := range schema.Views {
			view, ok := schemaState.views[view.Name]
			if !ok {
				continue
			}
			if _, err := w.WriteString(getViewAnnouncement(view.name)); err != nil {
				return err
			}

			buf := &strings.Builder{}
			if err := view.toString(buf); err != nil {
				return err
			}
			if _, err := w.WriteString(buf.String()); err != nil {
				return err
			}
			delete(schemaState.views, view.name)
		}
	}
	return nil
}

func getTableAnnouncement(name string) string {
	return fmt.Sprintf("\n--\n-- Table structure for `%s`\n--\n", name)
}

func getViewAnnouncement(name string) string {
	return fmt.Sprintf("\n--\n-- View structure for `%s`\n--\n", name)
}
