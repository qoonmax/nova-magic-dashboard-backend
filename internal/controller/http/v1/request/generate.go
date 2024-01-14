package request

type GenerateRequest struct {
	Key               string            `json:"key"`
	ClientEnvironment ClientEnvironment `json:"client_environment"`
	Database          Database          `json:"database"`
}

type ClientEnvironment struct {
	PhpVersion         string `json:"php_version"`
	LaravelVersion     string `json:"laravel_version"`
	NovaVersion        string `json:"nova_version"`
	DatabaseDriverName string `json:"database_driver_name"`
}

type Database struct {
	Tables []Table `json:"tables"`
}

type Table struct {
	Name    string   `json:"name"`
	Columns []Column `json:"columns"`
	Data    []Data   `json:"data"`
}

type Column struct {
	Name          string `json:"name"`
	Type          string `json:"type"`
	Length        int    `json:"length,omitempty"`
	Unsigned      bool   `json:"unsigned,omitempty"`
	Nullable      bool   `json:"nullable"`
	AutoIncrement bool   `json:"auto_increment,omitempty"`
	PrimaryKey    bool   `json:"primary_key,omitempty"`
	Unique        bool   `json:"unique,omitempty"`
}

type Data struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	EmailVerifiedAt string `json:"email_verified_at,omitempty"`
	Password        string `json:"password"`
	RememberToken   string `json:"remember_token,omitempty"`
	CreatedAt       string `json:"created_at"`
	UpdatedAt       string `json:"updated_at"`
}
