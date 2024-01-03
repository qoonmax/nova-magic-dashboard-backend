package v1

type Request struct {
	LaravelVersion string `json:"laravel"`
	NovaVersion    string `json:"nova"`
	Key            string `json:"key"`
}
