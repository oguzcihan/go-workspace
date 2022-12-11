package main

import "ExampleApp/internal/core"

func main() {
	core.StartServer()

}

//type Manager struct {
//	FullName       string `json=full_name`
//	Position       string
//	Age            int32
//	YearsInCompany int32
//}
//
//func EncodeManager(manager *Manager) (io.Reader, error) {
//	u, err := json.Marshal(Manager{FullName: manager.FullName, Position: manager.Position, Age: manager.Age, YearsInCompany: manager.YearsInCompany})
//	if err != nil {
//		panic(err)
//	}
//	return bytes.NewReader(u), nil
//
//}
