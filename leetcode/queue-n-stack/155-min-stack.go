package main

type MinStack struct {
	data 	[]int
	mins 	[]int
}

func (this *MinStack) Push(x int)  {
	minsLength := len(this.mins)

	if minsLength == 0 || x <= this.mins[minsLength - 1] {
		this.mins = append(this.mins, x)
	}

	this.data = append(this.data, x)
}

func (this *MinStack) Pop()  {
	minsLength := len(this.mins)
	dataLength := len(this.data)

	if this.data[dataLength - 1] == this.mins[minsLength - 1] {
		this.mins = this.mins[:minsLength - 1]
	}

	this.data = this.data[:dataLength - 1]
}

func (this *MinStack) Top() int {
	dataLength := len(this.data)

	return this.data[dataLength - 1]
}

func (this *MinStack) GetMin() int {
	minsLength := len(this.mins)

	return this.mins[minsLength - 1]
}