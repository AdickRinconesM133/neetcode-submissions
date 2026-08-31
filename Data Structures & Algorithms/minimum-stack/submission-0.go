type MinStack struct {
	stack	[]int
	minStack	[]int
}

func Constructor() MinStack {
	return MinStack{
		stack: []int{},
		minStack: []int{},
	}
}

func (this *MinStack) Push(val int) {
	if len(this.stack) == 0 {
		this.stack = append(this.stack, val)
		this.minStack = append(this.minStack, val)	
		return
	} 

	this.stack = append(this.stack, val)

	if val <= this.GetMin() {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}

	topValue := this.Top()
	this.stack = this.stack[:len(this.stack)-1]

	if topValue == this.GetMin() {
		this.minStack = this.minStack[:len(this.minStack)-1]
	}
}

func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}
