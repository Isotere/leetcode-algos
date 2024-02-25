package unique_array

func removeDuplicates(nums []int) int {
	// Если элементов нет, либо их всего 1
	if len(nums) < 2 {
		return len(nums)
	}

	// Текущий "уникальный" элемент
	current := nums[0]
	// Позиция куда будет вставлен следующий уникальный элемент
	pos := 1

	for i := 1; i < len(nums); i++ {
		// Если текущий равен предыдущему
		if current == nums[i] {
			continue
		}

		// Если новое значение - обновляем сохраненное
		current = nums[i]

		// вставляем уникальный элемент на позицию, придуманную для него
		nums[pos] = current

		// следующая позиция для уникального элемента
		pos++
	}

	return pos
}
