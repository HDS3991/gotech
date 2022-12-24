package sql2gorm

type NullStyle int

const (
	NullDisable NullStyle = iota
	NullInSql
	NullInPointer
)

type Option func(*options)

type options struct {
	Charset        string
	Collation      string
	JsonTag        bool
	TablePrefix    string
	ColumnPrefix   string
	NotNullType    bool
	NullStyle      NullStyle
	Package        string
	GormType       bool
	ForceTableName bool
}

var defaultOptions = options{
	NullStyle: NullInPointer,
	Package:   "",
}

func WithCharset(charset string) Option {
	return func(o *options) {
		o.Charset = charset
	}
}

func WithCollation(collation string) Option {
	return func(o *options) {
		o.Collation = collation
	}
}

func WithJsonTag() Option {
	return func(o *options) {
		o.JsonTag = true
	}
}

func WithTablePrefix(tablePrefix string) Option {
	return func(o *options) {
		o.TablePrefix = tablePrefix
	}
}

func WithColumnPrefix(colPrefix string) Option {
	return func(o *options) {
		o.ColumnPrefix = colPrefix
	}
}

func WithNotNullType() Option {
	return func(o *options) {
		o.NotNullType = true
	}
}

func WithNullStyle(s NullStyle) Option {
	return func(o *options) {
		o.NullStyle = s
	}
}

func WithPackage(pkg string) Option {
	return func(o *options) {
		o.Package = pkg
	}
}

func WithGormType() Option {
	return func(o *options) {
		o.GormType = true
	}
}

func WithForceTableName() Option {
	return func(o *options) {
		o.ForceTableName = true
	}
}

func parseOption(options ...Option) options {
	o := defaultOptions
	for _, f := range options {
		f(&o)
	}
	if o.NotNullType {
		o.NullStyle = NullDisable
	}
	return o
}
