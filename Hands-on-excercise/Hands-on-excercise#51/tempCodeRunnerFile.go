for k, v := range xi {
		fmt.Println(k)
		for i, value := range v {
			fmt.Println(i, value)
		}
	}