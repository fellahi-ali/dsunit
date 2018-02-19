package dsunit
//
//import (
//	"fmt"
//	"github.com/viant/dsc"
//	"github.com/viant/toolbox"
//	"strings"
//)
//
//type DatastoreDatasetProvider struct {
//	dsc.Manager
//}
//
//func (p *DatastoreDatasetProvider) Get(table string, columns ...string) (*Dataset, error) {
//	if len(columns) == 0 {
//		columns = []string{"*"}
//	}
//	var rows = make([]*Record, 0)
//	sql := fmt.Sprintf("SELECT %v FROM %v ORDER BY 1", strings.Join(columns, ","), table)
//	err := p.Manager.ReadAllWithHandler(sql, nil, func(scanner dsc.Scanner) (toContinue bool, err error) {
//		row := &Record{
//			Values: make(map[string]interface{}),
//			Source: fmt.Sprintf("%v record: %v", sql, len(rows)+1),
//		}
//
//		rows = append(rows, row)
//		err = scanner.Scan(&row.Values)
//
//		keys := toolbox.MapKeysToStringSlice(row.Values)
//
//		for _, key := range keys {
//			if row.Values[key] == nil {
//				delete(row.Values, key)
//			}
//		}
//		return true, err
//	})
//	if err != nil {
//		return nil, err
//	}
//	var result = &Dataset{
//		TableDescriptor: &dsc.TableDescriptor{
//			Table: table,
//		},
//		Rows: rows,
//	}
//	return result, nil
//}
//
//func NewDatastoreDatasetProvider(manager dsc.Manager) *DatastoreDatasetProvider {
//	return &DatastoreDatasetProvider{
//		Manager: manager,
//	}
//}
