package list

import (
	mysql "github.com/ShkrutDenis/go-migrations/builder"
	"github.com/jmoiron/sqlx"
)

type CreateProductTable struct{}

func (m *CreateProductTable) GetName() string {
	return "CreateProductTable"
}

func (m *CreateProductTable) Up(con *sqlx.DB) {
	table := mysql.NewTable("profiles", con)
	table.Column("id").Type("int unsigned").Autoincrement()
	table.PrimaryKey("id")
	table.String("alamat", 500).Nullable()
	table.String("pekerjaan", 500).Nullable()
	table.String("nama_lengkap", 500).Nullable()
	table.String("pendidikan", 500).Nullable()
	table.String("nomor_telpon", 500).Nullable()
	table.Column("deleted_at").Type("datetime").Nullable()
	table.Column("user_id").Type("int unsigned")
	table.ForeignKey("user_id").
		Reference("users").
		On("id").
		OnDelete("cascade").
		OnUpdate("cascade")
	table.WithTimestamps()

	table.MustExec()
}

func (m *CreateProductTable) Down(con *sqlx.DB) {
	mysql.DropTable("products", con).MustExec()
}
