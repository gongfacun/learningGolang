package queue

type  Queue []int


func (q *Queue)Push(v int){

	*q=append(*q,v)
}

func (q *Queue)Pop()int  {
	tmp:=(*q)[0]
	*q=(*q)[1:]
	return tmp
}
func (q *Queue)IsEmpty()bool  {
	return len(*q)==0
}