package entity

type ElectricityInfo struct {
	CUSTOMERID      int8   `json:"CUSTOMERID"`
	DISC            string `json:"DISC"`
	HH              string `json:"HH"`
	CODE            string `json:"CODE"`
	TYPE            int    `json:"TYPE"`
	POWERSUPP       int8   `json:"POWERSUPP"`
	LINEID          int8   `json:"LINEID"`
	LINEID2         int8   `json:"LINEID2"`
	CUSTOMERS       int8   `json:"CUSTOMERS"`
	MEASUREMODE     int    `json:"MEASUREMODE"`
	ECONOMICPROTID  int8   `json:"ECONOMICPROTID"`
	DEMANDPROTID    int8   `json:"DEMANDPROTID"`
	BUSINESSCLASSID int8   `json:"BUSINESSCLASSID"`
	CAPABILITY      string `json:"CAPABILITY"`
	CREATETIME      string `json:"CREATETIME"`
	ADDRESS         string `json:"ADDRESS"`
	LINKMAN         string `json:"LINKMAN"`
	TELEPHONE       string `json:"TELEPHONE"`
	MOBILE          string `json:"MOBILE"`
	NOTE            string `json:"NOTE"`
	LASTTIME        string `json:"LASTTIME"`
}

type TaskInfo struct {
	ServiceID   int64  `json:"service_id"`
	ServiceName string `json:"service_name"`
	TaskID      int64  `json:"task_id"`
	HandleID    int64  `json:"handle_id"`
	DBName      string `json:"db_name"`
	TableName   string `json:"table_name"`
}
