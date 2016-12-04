package a_alg

type Comparable interface {
	LessThen(interface{}) bool
}

type node struct {
	link  *node
	value Comparable
}

type LinkedPriorityQueue struct {
	head *node
}

func (this *LinkedPriorityQueue) Add(item Comparable) {
	ref := this.head
	refOnRef := &this.head
	for ref != nil && ref.value.LessThen(item) {
		refOnRef = &ref.link
		ref = ref.link
	}
	*refOnRef = &node{link:ref, value:item}
}

func (this *LinkedPriorityQueue) Pop() (item Comparable) {
	if this.head != nil {
		item = this.head.value
		this.head = this.head.link
	}
	return
}