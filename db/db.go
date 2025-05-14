package db

func Put(service string, endpoint string) {
	services[service] = endpoint
}

func Get(service string) string {
	return services[service]
}

func init() {
	services = make(map[string]string)
}

var services map[string]string
