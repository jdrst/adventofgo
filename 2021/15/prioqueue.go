package main

type prioQueue []*node

func (pq prioQueue) Len() int { return len(pq) }

func (pq prioQueue) Less(i, j int) bool {
	return pq[i].weight < pq[j].weight
}

func (pq prioQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *prioQueue) Push(new interface{}) {
	n := len(*pq)
	item := new.(*node)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *prioQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}
