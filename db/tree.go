package db

type DatabaseNode struct {
	SchemaNodes    []SchemaNode
	ExtensionNodes []ExtensionNode
}

func (dn *DatabaseNode) Err() error {
	for _, schema := range dn.SchemaNodes {
		if err := schema.Err(); err != nil {
			return err
		}
	}
	return nil
}

type ExtensionNode struct {
	Extension Extension
}

type SchemaNode struct {
	Schema     Schema
	TableNodes []TableNode
}

func (sn *SchemaNode) Err() error {
	for _, tn := range sn.TableNodes {
		if err := tn.Err(); err != nil {
			return err
		}
	}
	return nil
}

type TableNode struct {
	Table       Table
	ColumnNodes []ColumnNode
	IndexNodes  []IndexNode
}

func (tn *TableNode) Err() error {
	for _, cn := range tn.ColumnNodes {
		if err := cn.Err(); err != nil {
			return err
		}
	}
	return nil
}

type IndexNode struct {
	Index Index
}
