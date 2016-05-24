package sort

func BubbleSort(data []int) []int {
    
    swap := -1
    aux := 0
	for swap !=0 {
        swap = 0
        for i := 0; i<len(data)-1; i++{
            if data[i] > data[i+1] {
                aux = data[i]
                data[i]=data[i+1]
                data[i+1]=aux
                swap++
            }
        }
	}

	return data
}
