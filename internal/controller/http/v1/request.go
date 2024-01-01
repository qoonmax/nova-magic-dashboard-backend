package v1

type Request struct {
	LaravelVersion string `json:"laravel_version"`
	NovaVersion    string `json:"nova_version"`
	Key            string `json:"key"`
}
