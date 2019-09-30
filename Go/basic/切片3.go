package main

import "fmt"

// 切片


func main() {
	
	var countryCapitaMap map[string]string // 创建集合

	countryCapitaMap = make(map[string]string)

	countryCapitaMap["France"] = "Paris"
	countryCapitaMap["Italy"] = "罗马"
	countryCapitaMap["Japan"] = "东京"
	countryCapitaMap["India"] = "新德里"

	 printCountry(countryCapitaMap)


	// 是否存在
	captial, ok := countryCapitaMap["美国"]
	if (ok) {
		fmt.Println("美国的首都是", captial)
	}else{
		fmt.Println("美国首都不存在")
	}


	fmt.Println("\n法国条目被删除\n")

	// delete
	delete(countryCapitaMap,"France")


	printCountry(countryCapitaMap)

}

func printCountry(countryCapitaMap map[string]string) {
	
	/*使用键输出地图值 */
	for country := range countryCapitaMap{
		fmt.Println(country,"首都是",countryCapitaMap[country])
	}

}


