package config

type Config struct{
	DataBase DataBaseConfig	`json:"database"`
}

type DataBaseConfig struct{
	SourcePath string	 `json:"source"`
	DriverName string	`json:"driver_name"`
}