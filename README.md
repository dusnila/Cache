# Cache

данный репазитория была создана для выполнения домашнего задания для курса GOLANG-NINJA

данная репазитория позволяет записывать значение в кэш в виде Key-value

	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId = cache.Get("userId")

	fmt.Println(userId)